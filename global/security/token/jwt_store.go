package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "orca-service/global/logger"
	"orca-service/global/security"
	"time"
)

type JwtStore struct {
	key []byte
}

type JWTClaims struct {
	security.UserDetail `json:"user_detail"`
	Roles               []string `json:"roles"`
	Permissions         []string `json:"permissions"`
	jwt.StandardClaims
}

func NewJwtStore(key []byte) *JwtStore {
	return &JwtStore{
		key,
	}
}

func (j *JwtStore) CreateAccessToken(user security.UserDetail) (string, error) {
	roles := user.Roles
	permissions := user.Permissions
	user.Password = ""
	user.Roles = nil
	user.Permissions = nil
	claims := JWTClaims{
		UserDetail:  user,
		Roles:       roles,
		Permissions: permissions,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600*24, // 过期时间 一天
			Issuer:    "orca",                      //签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.key)
}

// RefreshAccessToken refresh access token，由于jwt的token是无法撤销的，所以这里直接返回新的token
func (j *JwtStore) RefreshAccessToken(token string) (string, error) {
	user, err := j.VerifyAccessToken(token)
	if err != nil {
		return "", err
	}
	if user != nil {
		return j.CreateAccessToken(*user)
	}
	return "", errors.New("令牌不可用")
}

// RemoveAccessToken remove access token，jwt无法撤销token，所以这里不做任何操作
func (j *JwtStore) RemoveAccessToken(_ security.UserDetail) error {
	log.Debug("Remove access token")
	return nil
}

// VerifyAccessToken verify access token
func (j *JwtStore) VerifyAccessToken(token string) (*security.UserDetail, error) {
	c, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("令牌不可用")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("令牌过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("令牌无效")
			} else {
				return nil, fmt.Errorf("令牌不可用")
			}

		}
	}
	claims, ok := c.Claims.(*JWTClaims)
	if ok && c.Valid {
		// 组装UserDetail
		user := claims.UserDetail
		user.Roles = claims.Roles
		user.Permissions = claims.Permissions
		return &user, nil
	}
	return nil, nil
}
