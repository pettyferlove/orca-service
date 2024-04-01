package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orca-service/global"
	"orca-service/global/handler"
	"orca-service/global/model"
	"orca-service/global/util"
)

func AccessControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles := util.GetRoles(c)
		service := "orca-service"
		hasRole := false
		for _, userRole := range roles {
			enforce, err := global.Enforcer.Enforce(userRole, c.Request.URL.Path, c.Request.Method, service)
			if err != nil {
				break
			}
			if enforce {
				hasRole = true
				break
			}
		}
		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, model.Response{
				Code:       int(handler.PermissionNoAccess),
				Message:    "未授权访问",
				Successful: false,
			})
			return
		}
		c.Next()
	}
}
