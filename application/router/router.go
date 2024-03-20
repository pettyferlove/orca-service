package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orca-service/global/model"
	"time"
)

var (
	router = make([]func(v1 *gin.RouterGroup), 0)
)

func InitRouter(r *gin.Engine) {
	r.NoRoute(NoRouteHandler)
	r.GET("/ping", PingHandler)
	group := r.Group("/handler/v1")
	for _, f := range router {
		f(group)
	}
}

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, model.Response{
		Code:       1,
		Message:    "not found",
		Successful: false,
		Timestamp:  time.Now().Unix(),
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code:       0,
		Successful: true,
		Timestamp:  time.Now().Unix(),
	})
}
