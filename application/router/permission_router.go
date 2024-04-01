package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
)

func init() {
	router = append(router, registerPermissionRouter)

}

func registerPermissionRouter(group *gin.RouterGroup) {
	permission := handler.Permission{}
	permissionGroup := group.Group("/permissions")
	permissionGroup.GET("/page", permission.Page)
	permissionGroup.GET("/:id", permission.Get)
	permissionGroup.POST("/", permission.Create)
	permissionGroup.PUT("/:id", permission.Update)
	permissionGroup.DELETE("/:id", permission.Delete)
	permissionGroup.GET("/tree", permission.Tree)
	permissionGroup.GET("/menus/tree", permission.MenusTree)
	permissionGroup.PUT("/move/:id/:type", permission.Move)
}
