package service

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	log "orca-service/global/logger"
)

// Service 结构体用于处理请求
type Service struct {
	Errors   error
	DataBase *gorm.DB
	Redis    *redis.Client
}

// AddError 方法用于添加一个新的错误到 Service 结构体
func (service *Service) AddError(err error) {
	if service.Errors == nil {
		service.Errors = err
	} else if err != nil {
		log.Error("handler process error, error:%v", err)
		service.Errors = fmt.Errorf("%v; %w", service.Errors, err)
	}
}
