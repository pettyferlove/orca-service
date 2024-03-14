package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orca-service/global"
	"orca-service/global/entity"
)

func AccessControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("roles")
		tenantId := "000001"
		service := "orca-service"
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, entity.ResponseEntity{
				Code:       1,
				Message:    "User roles are missing. Please assign roles to the user.",
				Successful: false,
			})
			return
		}
		roles := value.([]string)

		hasRole := false
		for _, userRole := range roles {
			enforce, err := global.Enforcer.Enforce(tenantId, userRole, c.Request.URL.Path, c.Request.Method, service)
			if err != nil {
				break
			}
			if enforce {
				hasRole = true
				break
			}

		}
		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, entity.ResponseEntity{
				Code:       1,
				Message:    "You do not have the required roles to access this resource.",
				Successful: false,
			})
			return
		}
		c.Next()
	}
}
