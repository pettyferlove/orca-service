package router

import "github.com/gin-gonic/gin"

func init() {
	router = append(router, registerPermissionRouter)

}

func registerPermissionRouter(group *gin.RouterGroup) {
	helloGroup := group.Group("/permissions")
	helloGroup.GET("/page", nil)
	helloGroup.GET("/:id", nil)
	helloGroup.POST("/", nil)
	helloGroup.PUT("/:id", nil)
	helloGroup.DELETE("/:id", nil)
	helloGroup.GET("/tree", nil)
	helloGroup.GET("/menus/tree", nil)
	helloGroup.GET("/move/:id/:type", nil)
}
