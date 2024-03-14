package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	global "orca-service/global/entity"
)

var (
	router = make([]func(v1 *gin.RouterGroup), 0)
)

func InitRouter(r *gin.Engine) {
	group := r.Group("/api/v1")
	for _, f := range router {
		f(group)
	}
}

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, global.ResponseEntity{
		Code:       1,
		Message:    "not found",
		Successful: false,
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, global.ResponseEntity{
		Message: "pong",
	})
}
