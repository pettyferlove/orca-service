package api

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
	"orca-service/global/api"
	"orca-service/global/logger"
	"orca-service/global/security"
)

type Token struct {
	api.Api
}

type LoginInfo struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenInfo struct {
	Token string `json:"token"`
	Type  string `json:"type,default=Bearer"`
}

func (t Token) Create(c *gin.Context) {
	var loginInfo LoginInfo
	err := t.MakeContext(c).Bind(&loginInfo).Errors
	if err != nil {
		logger.Error(err.Error())
		t.ResponseBusinessError(1, err.Error())
		return
	}
	token := handler.Token{}
	userToken, err := token.Create(loginInfo.Username, loginInfo.Password)
	if err != nil {
		t.ResponseBusinessError(1, err.Error())
		return
	}
	t.Response(TokenInfo{Token: userToken, Type: "Bearer"})
	return
}

func (t Token) Delete(c *gin.Context) {
	t.MakeContext(c)
	// JWT Token无需删除，客户端扔掉即可，因为它是短期的，服务端不需记录它
	// 兼容后期Token加入Redis或者JWT Token加入Redis
	t.ResponseOk()
	return
}

// Refresh 刷新Token
func (t Token) Refresh(context *gin.Context) {
	t.MakeContext(context)
	originalClaims, exists := t.Context.Get("original_claims")
	if !exists {
		t.ResponseBusinessError(1, "Token parsing failed.")
		return
	} else {
		// 转换为JWTClaims
		var claims = originalClaims.(*security.JWTClaims)
		token := handler.Token{}
		// 刷新Token，传入Claims（解引用）
		userToken, err := token.Refresh(*claims)
		if err != nil {
			t.ResponseBusinessError(1, err.Error())
			return
		}
		t.Response(userToken)
		return
	}
	return
}
