package router

import (
	"github.com/gin-gonic/gin"
)

func init() {
	router = append(router, registerRoleRouter)

}

func registerRoleRouter(group *gin.RouterGroup) {
	helloGroup := group.Group("/roles")
	helloGroup.GET("/page", nil)
	helloGroup.GET("/:id", nil)
	helloGroup.POST("/", nil)
	helloGroup.PUT("/:id", nil)
	helloGroup.DELETE("/:id", nil)
	helloGroup.PUT("/:id/valid", nil)
	helloGroup.GET("/available/all", nil)
}
