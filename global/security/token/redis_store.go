package token

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"orca-service/global/security"
	"time"
)

var (
	defaultAuthToAccessKeyPrefix          = "security:authorization:auth_to_access:"
	defaultUsernameToAccessKeyPrefix      = "security:authorization:username_to_access:"
	defaultAbnormalAccessKeyPrefix        = "security:authorization:abnormal_access:"
	defaultAccessTokenValiditySeconds     = 60 * 60 * 24
	defaultAccessTokenRefreshCriticalTime = 60 * 60 * 2
)

type RedisStore struct {
	redis                          *redis.Client
	authToAccessKeyPrefix          string
	usernameToAccessKeyPrefix      string
	abnormalAccessKeyPrefix        string
	accessTokenValiditySeconds     int
	accessTokenRefreshCriticalTime int
}

type AuthenticationAbnormal struct {
	Message  string
	Username string
}

func NewRedisStore(redis *redis.Client) *RedisStore {
	return &RedisStore{
		redis:                          redis,
		authToAccessKeyPrefix:          defaultAuthToAccessKeyPrefix,
		usernameToAccessKeyPrefix:      defaultUsernameToAccessKeyPrefix,
		abnormalAccessKeyPrefix:        defaultAbnormalAccessKeyPrefix,
		accessTokenValiditySeconds:     defaultAccessTokenValiditySeconds,
		accessTokenRefreshCriticalTime: defaultAccessTokenRefreshCriticalTime,
	}
}

func (r *RedisStore) CreateAccessToken(user security.UserDetail) (string, error) {
	username := user.Username
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("create token failed")
	}
	token := uuidObj.String()
	serializedData, _ := json.Marshal(user)
	err = r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, token), string(serializedData), time.Duration(r.accessTokenValiditySeconds)*time.Second).Err()
	if err != nil {
		return "", errors.New("create token failed")
	}
	// 判断是否单点登录
	if !AllowMultiPoint {
		oldToken, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, username)).Result()
		if err == nil && oldToken != "" {
			r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, oldToken))
			r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, username))
			// 记录旧的token异常登录
			abnormalObj := AuthenticationAbnormal{Message: "The account is logged in elsewhere and you have been forced offline", Username: username}
			abnormalData, _ := json.Marshal(abnormalObj)
			r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", r.abnormalAccessKeyPrefix, oldToken), string(abnormalData), time.Duration(r.accessTokenValiditySeconds)*time.Second)
		}
		// 将用户名和token绑定，用户检测重复登录
		r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, username), token, time.Duration(r.accessTokenValiditySeconds)*time.Second)
	}
	return token, nil
}

func (r *RedisStore) RefreshAccessToken(token string) (string, error) {
	userDetails, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, token)).Result()
	if err != nil {
		return "", errors.New("令牌无效")
	}

	var userDetailsObj security.UserDetail
	err = json.Unmarshal([]byte(userDetails), &userDetailsObj)
	if err != nil {
		return "", errors.New("令牌无效")
	}

	serializedData, _ := json.Marshal(userDetailsObj)
	err = r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, token), string(serializedData), time.Duration(r.accessTokenValiditySeconds)*time.Second).Err()
	if err != nil {
		return "", errors.New("令牌无效")
	}

	err = r.redis.LPush(context.Background(), fmt.Sprintf("%s:%s", r.usernameToAccessKeyPrefix, userDetailsObj.Username), token).Err()
	if err != nil {
		return "", errors.New("令牌无效")
	}

	err = r.redis.Expire(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, userDetailsObj.Username), time.Duration(r.accessTokenValiditySeconds)*time.Second).Err()
	return token, errors.New("令牌无效")
}

func (r *RedisStore) RemoveAccessToken(user security.UserDetail) error {
	keys, err := r.redis.LRange(context.Background(), fmt.Sprintf("%s:%s", r.usernameToAccessKeyPrefix, user.Username), 0, -1).Result()
	if err != nil {
		return errors.New("令牌无效")
	}
	for _, key := range keys {
		err = r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", r.abnormalAccessKeyPrefix, key)).Err()
		if err != nil {
			return errors.New("令牌无效")
		}
	}
	err = r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", r.usernameToAccessKeyPrefix, user.Username)).Err()
	return errors.New("令牌无效")
}

func (r *RedisStore) VerifyAccessToken(token string) (*security.UserDetail, error) {
	duration, err := r.redis.TTL(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, token)).Result()
	if err != nil || duration.Seconds() <= 0 {
		return nil, errors.New("令牌无效")
	}
	userDetails, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", r.authToAccessKeyPrefix, token)).Result()
	if err != nil {
		return nil, errors.New("令牌无效")
	}
	// 数据反序列化为UserDetails结构体
	var userDetailsObj security.UserDetail
	err = json.Unmarshal([]byte(userDetails), &userDetailsObj)
	if err != nil {
		return nil, errors.New("令牌无效")
	}
	// 处理异常登录
	if !AllowMultiPoint {
		authAbnormal, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", r.abnormalAccessKeyPrefix, token)).Result()
		if err == nil {
			var authAbnormalObj AuthenticationAbnormal
			err = json.Unmarshal([]byte(authAbnormal), &authAbnormalObj)
			if err == nil {
				r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", r.abnormalAccessKeyPrefix, token))
				return nil, fmt.Errorf(authAbnormalObj.Message)
			}
		}
	}
	if duration < time.Duration(r.accessTokenRefreshCriticalTime)*time.Second {
		_, err := r.RefreshAccessToken(token)
		if err != nil {
			return nil, err
		}
	}
	return &userDetailsObj, nil

}
