package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"orca-service/application/service"
	"orca-service/global"
	"orca-service/global/handler"
	"orca-service/global/logger"
	"orca-service/global/security"
	"orca-service/global/security/token"
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
	var userService = service.User{}
	err := t.MakeContext(c).MakeService(&userService.Service).Bind(&loginRequest).Errors
	if err != nil {
		logger.Error(err.Error())
		t.ResponseInvalidArgument(err.Error())
		return
	}
	userDetail := userService.LoadUserByUsername(loginRequest.Username)
	if userService.Errors != nil {
		logger.Error(userService.Errors.Error())
		t.ResponseUnauthorizedMessage("User does not exist")
		return
	}
	// 使用BCrypt进行密码校验
	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(loginRequest.Password))
	// 清空密码
	userDetail.Password = ""
	if err != nil {
		t.ResponseUnauthorizedMessage("Password error")
		return
	}
	store := token.NewRedisStore(global.RedisClient)
	accessToken, err := store.CreateAccessToken(*userDetail)
	if err != nil {
		return

	}
	if err != nil {
		t.ResponseUnauthorizedMessage("Token creation failed")
		return
	}
	t.Response(LoginResponse{Token: accessToken, Type: "Bearer"})
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
		t.ResponseBusinessError(1, "Token parsing failed")
		return
	} else {
		// 转换为JWTClaims
		var _ = originalClaims.(*security.JWTClaims)

		t.Response("")
		return
	}
}
