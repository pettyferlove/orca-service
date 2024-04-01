package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
	"orca-service/global/security/middleware"
)

func init() {
	router = append(router, registerHelloRouter)

}

func registerHelloRouter(group *gin.RouterGroup) {
	hello := handler.Hello{}
	helloGroup := group.Group("/hello",
		middleware.AuthenticationMiddleware(),
		middleware.AccessControlMiddleware(),
	)
	helloGroup.GET("/page", hello.Page)
	helloGroup.GET("/:id", hello.Get)
	helloGroup.POST("/", hello.Post)
	helloGroup.PUT("/:id", hello.Put)
	helloGroup.DELETE("/:id", hello.Delete)
}
