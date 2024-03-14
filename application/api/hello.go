package api

import (
	"github.com/gin-gonic/gin"
	"orca-service/global/api"
	"orca-service/global/log"
)

type Hello struct {
	api.Api
}

func (h Hello) Page(c *gin.Context) {
	h.MakeContext(c)
	log.Info("hello page")
	h.ResponseOk()
	return
}

func (h Hello) Get(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	//global.Enforcer.AddNamedPolicies("p", [][]string{
	//	{"000001", "ROLE_USER", "/api/v1/hello/page", "GET", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/api/v1/hello/:id", "GET", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/api/v1/hello", "POST", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/api/v1/hello/:id", "PUT", "orca-service", "allow"},
	//	{"000001", "ROLE_USER", "/api/v1/hello/:id", "DELETE", "orca-service", "allow"},
	//})
	log.Info("hello get id:" + id)
	h.ResponseOk()
	return
}

func (h Hello) Post(c *gin.Context) {
	h.MakeContext(c)
	log.Info("hello post")
	h.ResponseOk()
	return
}

func (h Hello) Put(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	log.Info("hello put, id:" + id)
	h.ResponseOk()
	return
}

func (h Hello) Delete(c *gin.Context) {
	h.MakeContext(c)
	id := c.Param("id")
	log.Info("hello delete, id:" + id)
	h.ResponseOk()
	return
}
