package router

import "github.com/gin-gonic/gin"

func init() {
	router = append(router, registerUserRouter)

}

func registerUserRouter(group *gin.RouterGroup) {
	helloGroup := group.Group("/users")
	helloGroup.GET("/page", nil)
	helloGroup.GET("/:id", nil)
	helloGroup.POST("/", nil)
	helloGroup.PUT("/:id", nil)
	helloGroup.DELETE("/:id", nil)
	helloGroup.PUT("/current/roles", nil)
	helloGroup.GET("/:id/roles", nil)
	helloGroup.PUT("/:id/roles", nil)
}
