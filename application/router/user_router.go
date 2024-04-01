package router

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/handler"
)

func init() {
	router = append(router, registerUserRouter)

}

func registerUserRouter(group *gin.RouterGroup) {
	user := handler.User{}
	userGroup := group.Group("/users")
	userGroup.GET("/page", user.Page)
	userGroup.GET("/:id", user.Get)
	userGroup.POST("/", user.Create)
	userGroup.PUT("/:id", user.Update)
	userGroup.DELETE("/:id", user.Delete)
	userGroup.GET("/current/roles", user.CurrentRoles)
	userGroup.GET("/:id/roles", user.GetRoles)
	userGroup.PUT("/:id/roles", user.UpdateRoles)
}
