package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"net/http"
	"orca-service/global"
	log "orca-service/global/logger"
	"orca-service/global/model"
	"orca-service/global/util"
)

// Api 结构体用于处理 API 请求和响应
type Api struct {
	Context  *gin.Context
	Errors   error
	DataBase *gorm.DB
}

// AddError 方法用于添加一个新的错误到 Api 结构体
func (api *Api) AddError(err error) {
	if api.Errors == nil {
		api.Errors = err
	} else if err != nil {
		log.Error("api process error, error:%v", err)
		api.Errors = fmt.Errorf("%v; %w", api.Errors, err)
	}
}

// MakeContext 方法用于设置 Api 结构体的上下文
func (api *Api) MakeContext(c *gin.Context) *Api {
	api.Context = c
	api.DataBase = global.DatabaseClient.WithContext(c)
	return api
}

// Bind 方法用于绑定请求数据到指定的结构体
func (api *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = cache.GetBinding(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = api.Context.ShouldBindUri(d)
		} else {
			err = api.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			api.AddError(errors.New("request body is not present anymore. "))
			break
		}
		if err != nil {
			api.AddError(err)
			break
		}
	}
	if err == nil {
		validate, s := util.Validate(d)
		if !validate {
			api.AddError(errors.New(s))
		}
	}
	return api
}

// Response 方法用于发送一个成功的响应
func (api *Api) Response(object any) {
	api.Context.JSON(http.StatusOK, model.Response{
		Code:       0,
		Data:       object,
		Successful: true,
	})
}

// ResponseOk 方法用于发送一个成功的响应，但没有数据返回
func (api *Api) ResponseOk() {
	api.Context.JSON(http.StatusOK, model.Response{
		Code:       0,
		Data:       nil,
		Successful: true,
	})
}

// ResponseMessage 方法用于发送一个成功的响应，但没有数据返回
func (api *Api) ResponseMessage(message string) {
	api.Context.JSON(http.StatusOK, model.Response{
		Code:       0,
		Message:    message,
		Data:       nil,
		Successful: true,
	})
}

// ResponseBusinessError 方法用于发送一个业务错误的响应
func (api *Api) ResponseBusinessError(code int, message string) {
	api.Context.AbortWithStatusJSON(http.StatusOK, model.Response{
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseError 方法用于发送一个错误的响应
func (api *Api) ResponseError(code int, message string) {
	api.Context.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}
