package handler

import (
	"github.com/dgrijalva/jwt-go"
	log "orca-service/global/logger"
	"orca-service/global/security"
	"orca-service/global/security/model"
	"time"
)

type Token struct{}

func (t *Token) Create(username string, password string) (string, error) {
	log.Debug("username: %s, password: %s", username, password)
	j := security.NewJWT()
	token, err := j.CreateToken(security.JWTClaims{
		UserDetail: model.UserDetail{
			Id:       "0000001",
			UserName: "Alex Pettyfer",
			NickName: "Pettyfer",
			Avatar:   "https://avatars3.githubusercontent.com/u/42489393?s=460&v=4",
			Email:    "pettyferlove@live.cn",
		},
		Roles: []string{
			"ROLE_ADMIN",
			"ROLE_USER",
		},
		Permissions: []string{
			"sys:user:*",
			"api:hello:get",
		},
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600*24, // 过期时间 一天
			Issuer:    "orca",                      //签名的发行者
		},
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *Token) Refresh(claims security.JWTClaims) (string, error) {
	j := security.NewJWT()
	token, err := j.CreateToken(security.JWTClaims{
		UserDetail: claims.UserDetail,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600*24, // 过期时间 一天
			Issuer:    "orca",                      //签名的发行者
		},
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
