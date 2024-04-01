package handler

import (
	"github.com/gin-gonic/gin"
	"orca-service/global/handler"
)

type User struct {
	handler.Handler
}

func (u User) Page(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) Get(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) Create(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) Update(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) Delete(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) CurrentRoles(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) GetRoles(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}

func (u User) UpdateRoles(c *gin.Context) {
	u.MakeContext(c)
	u.ResponseOk()
	return
}
