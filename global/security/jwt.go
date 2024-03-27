package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	// 加载私钥
	return &JWT{
		[]byte("fpxt@GeZNUErvj!DXb7XMyeP_Mezhae9"),
	}
}

type JWTClaims struct {
	UserDetail  `json:"user_detail"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParserToken(token string) (*JWTClaims, error) {
	c, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token is not available")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token expires")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("invalid token")
			} else {
				return nil, fmt.Errorf("token is not available")
			}

		}
	}
	if claims, ok := c.Claims.(*JWTClaims); ok && c.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token is invalid")
}
