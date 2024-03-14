package common

import (
	"orca-service/global"
	"orca-service/global/config"
	"testing"
)

func TestInitRedis(t *testing.T) {
	global.Config.Redis = config.Redis{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
		Database: 0,
		Pool:     config.RedisConnectionPool{PoolSize: 200, MinIdle: 50},
	}
	err := InitRedis()
	if err != nil {
		t.Errorf("failed to initialize redis: %v", err)
	}
}
