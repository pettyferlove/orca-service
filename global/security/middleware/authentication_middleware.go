package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"orca-service/global/model"
	"orca-service/global/security"
	"orca-service/global/util"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Code:       1,
				Message:    "The request does not carry a token and there is no access permission.",
				Data:       nil,
				Successful: false,
			})
			return
		}
		j := security.NewJWT()
		jwtClaims, err := j.ParserToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				Code:       1,
				Message:    "Token parsing failed.",
				Data:       nil,
				Successful: false,
			})
			return
		}
		// 清除StandardClaims
		jwtClaims.StandardClaims = jwt.StandardClaims{}
		c.Set("original_claims", jwtClaims)
		c.Set("user_detail", jwtClaims.UserDetail)
		c.Set("roles", jwtClaims.Roles)
		c.Set("permissions", jwtClaims.Permissions)
		util.WithContext(c, jwtClaims)
		c.Next()
	}
}
