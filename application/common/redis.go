package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"orca-service/global"
	"orca-service/global/log"
)

var ctx = context.Background()

func InitRedis() error {

	redisConfig := global.Config.Redis
	pool := redisConfig.Pool

	address := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)

	r := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     redisConfig.Password,
		DB:           redisConfig.Database,
		PoolSize:     pool.PoolSize,
		MinIdleConns: pool.MinIdle,
	})
	pong, err := r.Ping(ctx).Result()
	if err != nil {
		log.Error("redis ping failed: %v", err)
		return err
	} else {
		log.Info("redis ping response: %s", pong)
	}
	global.RedisClient = r
	return nil
}
