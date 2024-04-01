package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
	"orca-service/global/security/middleware"
)

func init() {
	router = append(router, registerTokenRouter)
}

func registerTokenRouter(group *gin.RouterGroup) {
	token := handler.Token{}
	group.POST("/tokens", token.Create)
	group.DELETE("/tokens", middleware.AuthenticationMiddleware(), token.Delete)
	group.POST("/tokens/refresh", middleware.AuthenticationMiddleware(), token.Refresh)
}
