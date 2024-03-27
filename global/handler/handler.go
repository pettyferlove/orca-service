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
	"orca-service/global/service"
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
func (handler *Handler) AddError(err error) {
	if handler.Errors == nil {
		handler.Errors = err
	} else if err != nil {
		log.Error("handler process error, error:%v", err)
		handler.Errors = fmt.Errorf("%v; %w", handler.Errors, err)
	}
}

// MakeContext 方法用于设置 Handler 结构体的上下文
func (handler *Handler) MakeContext(c *gin.Context) *Handler {
	handler.Context = c
	handler.DataBase = global.DatabaseClient.WithContext(c)
	handler.Redis = global.RedisClient
	return handler
}

// MakeService 方法用于设置 Handler 结构体的服务
func (handler *Handler) MakeService(service *service.Service) *Handler {
	service.DataBase = handler.DataBase
	service.Redis = handler.Redis
	return handler
}

// Bind 方法用于绑定请求数据到指定的结构体
func (handler *Handler) Bind(d interface{}, bindings ...binding.Binding) *Handler {
	var err error
	if len(bindings) == 0 {
		bindings = cache.GetBinding(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = handler.Context.ShouldBindUri(d)
		} else {
			err = handler.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			handler.AddError(errors.New("request body is not present anymore. "))
			break
		}
		if err != nil {
			handler.AddError(err)
			break
		}
	}
	if err == nil {
		validate, s := util.Validate(d)
		if !validate {
			handler.AddError(errors.New(s))
		}
	}
	return handler
}

// Response 方法用于发送一个成功的响应
func (handler *Handler) Response(object any) {
	handler.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(Success),
		Data:       object,
		Successful: true,
	})
}

// ResponseOk 方法用于发送一个成功的响应，但没有数据返回
func (handler *Handler) ResponseOk() {
	handler.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(Success),
		Data:       nil,
		Successful: true,
	})
}

// ResponseMessage 方法用于发送一个成功的响应，但没有数据返回
func (handler *Handler) ResponseMessage(message string) {
	handler.Context.JSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(Success),
		Message:    message,
		Data:       nil,
		Successful: true,
	})
}

// ResponseBusinessError 方法用于发送一个业务错误的响应
func (handler *Handler) ResponseBusinessError(code int, message string) {
	handler.Context.AbortWithStatusJSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseError 方法用于发送一个错误的响应
func (handler *Handler) ResponseError(code int, message string) {
	handler.Context.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       code,
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseInvalidArgument 方法用于发送一个无效参数的响应
func (handler *Handler) ResponseInvalidArgument(message string) {
	handler.Context.AbortWithStatusJSON(http.StatusOK, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(InvalidArgument),
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseNotFound 方法用于发送一个未找到的响应
func (handler *Handler) ResponseNotFound() {
	handler.Context.AbortWithStatusJSON(http.StatusNotFound, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(Failure),
		Message:    "Not found",
		Data:       nil,
		Successful: false,
	})
}

// ResponseUnauthorized 方法用于发送一个未授权的响应
func (handler *Handler) ResponseUnauthorized() {
	handler.Context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(UserAuthenticateError),
		Message:    "Unauthorized",
		Data:       nil,
		Successful: false,
	})
}

// ResponseUnauthorizedMessage 方法用于发送一个未授权的响应
func (handler *Handler) ResponseUnauthorizedMessage(message string) {
	handler.Context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(UserAuthenticateError),
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseForbidden 方法用于发送一个禁止访问的响应
func (handler *Handler) ResponseForbidden() {
	handler.Context.AbortWithStatusJSON(http.StatusForbidden, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(PermissionNoAccess),
		Message:    "Forbidden",
		Data:       nil,
		Successful: false,
	})
}

// ResponseForbiddenMessage 方法用于发送一个禁止访问的响应
func (handler *Handler) ResponseForbiddenMessage(message string) {
	handler.Context.AbortWithStatusJSON(http.StatusForbidden, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(PermissionNoAccess),
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}

// ResponseBadRequest 方法用于发送一个错误的请求的响应
func (handler *Handler) ResponseBadRequest() {
	handler.Context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(ServiceException),
		Message:    "Bad request",
		Data:       nil,
		Successful: false,
	})
}

// ResponseBadRequestMessage 方法用于发送一个错误的请求的响应
func (handler *Handler) ResponseBadRequestMessage(message string) {
	handler.Context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Timestamp:  time.Now().Unix(),
		Code:       int(ServiceException),
		Message:    message,
		Data:       nil,
		Successful: false,
	})
}
