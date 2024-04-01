package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"orca-service/global/config"
)

const (
	// ConfigFilePath 配置文件路径
	ConfigFilePath = "./config/config.yaml"
)

// 全局变量
var (
	// Config 配置文件
	Config config.Config
	// DatabaseClient MySQL客户端
	DatabaseClient *gorm.DB
	RedisClient    *redis.Client
	Enforcer       *casbin.SyncedEnforcer
)
