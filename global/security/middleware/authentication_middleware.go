package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orca-service/global"
	"orca-service/global/handler"
	"orca-service/global/model"
	store "orca-service/global/security/token"
	"orca-service/global/util"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Code:       int(handler.UserAuthenticateError),
				Message:    "无访问权限",
				Data:       nil,
				Successful: false,
			})
			return
		}
		s := store.NewRedisStore(global.RedisClient)
		detail, err := s.VerifyAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Code:       int(handler.UserAuthenticateError),
				Message:    err.Error(),
				Data:       nil,
				Successful: false,
			})
			return
		}
		c.Set(util.UserDetailKey, detail)
		c.Set(util.AccessTokenKey, token)
		c.Next()
	}
}
