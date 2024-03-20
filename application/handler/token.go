package handler

import (
	"github.com/gin-gonic/gin"
	"orca-service/global/handler"
	"orca-service/global/logger"
	"orca-service/global/security"
)

type Token struct {
	handler.Handler
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type,default=Bearer"`
}

func (t Token) Create(c *gin.Context) {
	var loginRequest LoginRequest
	err := t.MakeContext(c).Bind(&loginRequest).Errors
	if err != nil {
		logger.Error(err.Error())
		t.ResponseBusinessError(1, err.Error())
		return
	}
	if err != nil {
		t.ResponseBusinessError(1, err.Error())
		return
	}
	t.Response(LoginResponse{Token: "", Type: "Bearer"})
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
		var _ = originalClaims.(*security.JWTClaims)

		t.Response("")
		return
	}
	return
}
