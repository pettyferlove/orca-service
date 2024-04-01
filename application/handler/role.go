package handler

import (
	"github.com/gin-gonic/gin"
	"orca-service/global/handler"
)

type Role struct {
	handler.Handler
}

func (r Role) Page(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) Get(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) Create(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) Update(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) Delete(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) Valid(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}

func (r Role) AvailableAll(c *gin.Context) {
	r.MakeContext(c)
	r.ResponseOk()
	return
}
