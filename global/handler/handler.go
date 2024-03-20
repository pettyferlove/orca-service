package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"net/http"
	"orca-service/global"
	log "orca-service/global/logger"
	"orca-service/global/model"
	"orca-service/global/util"
	"time"
)

// Handler 结构体用于处理请求
type Handler struct {
	Context  *gin.Context
	Errors   error
	DataBase *gorm.DB
	Redis    *redis.Client
}

// AddError 方法用于添加一个新的错误到 Handler 结构体
func (api *Handler) AddError(err error) {
	if api.Errors == nil {
		api.Errors = err
	} else if err != nil {
		log.Error("handler process error, error:%v", err)
		api.Errors = fmt.Errorf("%v; %w", api.Errors, err)
	}
}

// MakeContext 方法用于设置 Handler 结构体的上下文
func (api *Handler) MakeContext(c *gin.Context) *Handler {
	api.Context = c
	api.DataBase = global.DatabaseClient.WithContext(c)
	api.Redis = global.RedisClient
	return api
}

// Bind 方法用于绑定请求数据到指定的结构体
func (api *Handler) Bind(d interface{}, bindings ...binding.Binding) *Handler {
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
func (api *Handler) Response(object any) {
	api.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       0,
		Data:       object,
		Successful: true,
	})
}

// ResponseOk 方法用于发送一个成功的响应，但没有数据返回
func (api *Handler) ResponseOk() {
	api.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       0,
		Data:       nil,
		Successful: true,
	})
}

// ResponseMessage 方法用于发送一个成功的响应，但没有数据返回
func (api *Handler) ResponseMessage(message string) {
	api.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       0,
		Message:    message,
		Data:       nil,
		Successful: true,
	})
}

// ResponseBusinessError 方法用于发送一个业务错误的响应
func (api *Handler) ResponseBusinessError(code int, message string) {
	api.Context.AbortWithStatusJSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseError 方法用于发送一个错误的响应
func (api *Handler) ResponseError(code int, message string) {
	api.Context.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseNotFound 方法用于发送一个未找到的响应
func (api *Handler) ResponseNotFound() {
	api.Context.AbortWithStatusJSON(http.StatusNotFound, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       1,
		Message:    "not found",
		Data:       nil,
		Successful: false,
	})
}

// ResponseUnauthorized 方法用于发送一个未授权的响应
func (api *Handler) ResponseUnauthorized() {
	api.Context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       2001,
		Message:    "unauthorized",
		Data:       nil,
		Successful: false,
	})
}

// ResponseUnauthorizedMessage 方法用于发送一个未授权的响应
func (api *Handler) ResponseUnauthorizedMessage(message string) {
	api.Context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       2001,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseForbidden 方法用于发送一个禁止访问的响应
func (api *Handler) ResponseForbidden() {
	api.Context.AbortWithStatusJSON(http.StatusForbidden, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       4001,
		Message:    "forbidden",
		Data:       nil,
		Successful: false,
	})
}

// ResponseForbiddenMessage 方法用于发送一个禁止访问的响应
func (api *Handler) ResponseForbiddenMessage(message string) {
	api.Context.AbortWithStatusJSON(http.StatusForbidden, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       4001,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseBadRequest 方法用于发送一个错误的请求的响应
func (api *Handler) ResponseBadRequest() {
	api.Context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       9999,
		Message:    "bad request",
		Data:       nil,
		Successful: false,
	})
}

// ResponseBadRequestMessage 方法用于发送一个错误的请求的响应
func (api *Handler) ResponseBadRequestMessage(message string) {
	api.Context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       9999,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}
