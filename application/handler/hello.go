package handler

import (
	"github.com/gin-gonic/gin"
	"orca-service/global/handler"
	"orca-service/global/logger"
)

type Hello struct {
	handler.Handler
}

func (h Hello) Page(c *gin.Context) {
	h.MakeContext(c)
	logger.Info("hello page")
	h.ResponseOk()
	return
}

func (h Hello) Get(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	//global.Enforcer.AddNamedPolicies("p", [][]string{
	//	{"000001", "ROLE_USER", "/handler/v1/hello/page", "GET", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/handler/v1/hello/:id", "GET", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/handler/v1/hello", "POST", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/handler/v1/hello/:id", "PUT", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/handler/v1/hello/:id", "DELETE", "orca-service", "allow"},
	//})
	logger.Info("hello get id:" + id)
	h.ResponseOk()
	return
}

func (h Hello) Post(c *gin.Context) {
	h.MakeContext(c)
	logger.Info("hello post")
	h.ResponseOk()
	return
}

func (h Hello) Put(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	logger.Info("hello put, id:" + id)
	h.ResponseOk()
	return
}

func (h Hello) Delete(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	logger.Info("hello delete, id:" + id)
	h.ResponseOk()
	return
}
