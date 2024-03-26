package token

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"orca-service/global/security/model"
	"time"
)

var (
	authToAccessKeyPrefix          = "security:authorization:auth_to_access:"
	usernameToAccessKeyPrefix      = "security:authorization:username_to_access:"
	abnormalAccessKeyPrefix        = "security:authorization:abnormal_access:"
	accessTokenValiditySeconds     = 60 * 60 * 24
	accessTokenRefreshCriticalTime = 60 * 60 * 2
)

type RedisStore struct {
	redis *redis.Client
}

type AuthenticationAbnormal struct {
	Message  string
	Username string
}

func NewRedisStore(redis *redis.Client) *RedisStore {
	return &RedisStore{
		redis: redis,
	}
}

func (r RedisStore) CreateAccessToken(user model.UserDetail) (string, error) {
	username := user.Username

	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	token := uuidObj.String()

	serializedData, _ := json.Marshal(user)
	err = r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, token), string(serializedData), time.Duration(accessTokenValiditySeconds)*time.Second).Err()
	if err != nil {
		return "", err
	}
	// 判断是否单点登录
	if !AllowMultiPoint {
		oldToken, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, username)).Result()
		if err == nil && oldToken != "" {
			r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, oldToken))
			r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, username))

			// 记录旧的token异常登录
			abnormalObj := AuthenticationAbnormal{Message: "帐号在其他地方登录，您已被迫下线", Username: username}
			abnormalData, _ := json.Marshal(abnormalObj)
			r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", abnormalAccessKeyPrefix, oldToken), string(abnormalData), time.Duration(accessTokenValiditySeconds)*time.Second)
		}
		// 将用户名和token绑定，用户检测重复登录
		r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, username), token, time.Duration(accessTokenValiditySeconds)*time.Second)
	}

	return token, nil
}

func (r RedisStore) RefreshAccessToken(token string) (string, error) {
	userDetails, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, token)).Result()
	if err != nil {
		return "", err
	}

	var userDetailsObj model.UserDetail
	err = json.Unmarshal([]byte(userDetails), &userDetailsObj)
	if err != nil {
		return "", err
	}

	serializedData, _ := json.Marshal(userDetailsObj)
	err = r.redis.Set(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, token), string(serializedData), time.Duration(accessTokenValiditySeconds)*time.Second).Err()
	if err != nil {
		return "", err
	}

	err = r.redis.LPush(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, userDetailsObj.Username), token).Err()
	if err != nil {
		return "", err
	}

	err = r.redis.Expire(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, userDetailsObj.Username), time.Duration(accessTokenValiditySeconds)*time.Second).Err()
	return "", err
}

func (r RedisStore) RemoveAccessToken(user model.UserDetail) error {
	keys, err := r.redis.LRange(context.Background(), fmt.Sprintf("%s:%s", usernameToAccessKeyPrefix, user.Username), 0, -1).Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		err = r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", abnormalAccessKeyPrefix, key)).Err()
		if err != nil {
			return err
		}
	}
	err = r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", usernameToAccessKeyPrefix, user.Username)).Err()
	return err
}

func (r RedisStore) VerifyAccessToken(token string) (*model.UserDetail, error) {
	duration, err := r.redis.TTL(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, token)).Result()
	if err != nil || duration.Seconds() <= 0 {
		return nil, err
	}
	userDetails, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", authToAccessKeyPrefix, token)).Result()
	if err != nil {
		return nil, err
	}
	// 数据反序列化为UserDetails结构体
	var userDetailsObj model.UserDetail
	err = json.Unmarshal([]byte(userDetails), &userDetailsObj)
	if err != nil {
		return nil, err
	}
	// 处理异常登录
	if !AllowMultiPoint {
		authAbnormal, err := r.redis.Get(context.Background(), fmt.Sprintf("%s:%s", abnormalAccessKeyPrefix, token)).Result()
		if err == nil {
			var authAbnormalObj AuthenticationAbnormal
			err = json.Unmarshal([]byte(authAbnormal), &authAbnormalObj)
			if err == nil {
				r.redis.Del(context.Background(), fmt.Sprintf("%s:%s", abnormalAccessKeyPrefix, token))
				return nil, fmt.Errorf(authAbnormalObj.Message)
			}
		}
	}
	if duration < time.Duration(accessTokenRefreshCriticalTime)*time.Second {
		_, err := r.RefreshAccessToken(token)
		if err != nil {
			return nil, err
		}
	}
	return &userDetailsObj, nil

}
