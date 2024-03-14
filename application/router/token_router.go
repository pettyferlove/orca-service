package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/api"
	"orca-service/global/security/middleware"
)

func init() {
	router = append(router, registerTokenRouter)
}

func registerTokenRouter(group *gin.RouterGroup) {
	token := api.Token{}
	group.POST("/tokens", token.Create)
	group.DELETE("/tokens", token.Delete)
	group.POST("/tokens/refresh",
		middleware.AuthenticationMiddleware(),
		token.Refresh,
	)
}
