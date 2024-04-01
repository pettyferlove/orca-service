package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
)

func init() {
	router = append(router, registerRoleRouter)
}

func registerRoleRouter(group *gin.RouterGroup) {
	role := handler.Role{}
	roleGroup := group.Group("/roles")
	roleGroup.GET("/page", role.Page)
	roleGroup.GET("/:id", role.Get)
	roleGroup.POST("/", role.Create)
	roleGroup.PUT("/:id", role.Update)
	roleGroup.DELETE("/:id", role.Delete)
	roleGroup.PUT("/:id/valid", role.Valid)
	roleGroup.GET("/available/all", role.AvailableAll)
}
