package handler

import (
	"github.com/gin-gonic/gin"
	"orca-service/application/model"
	"orca-service/global/handler"
)

type Permission struct {
	handler.Handler
}

func (p Permission) Page(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) Get(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) Create(c *gin.Context) {
	var permissionCreate model.PermissionCreate
	err := p.MakeContext(c).Bind(&permissionCreate).Errors
	if err != nil {
		p.ResponseInvalidArgument(err.Error())
		return
	}
	p.ResponseOk()
	return
}

func (p Permission) Update(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) Delete(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) Tree(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) MenusTree(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}

func (p Permission) Move(c *gin.Context) {
	p.MakeContext(c)
	p.ResponseOk()
	return
}
