package cache

import (
	"bitstorm/configs"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
)

var (
	redisConn *redis.Client
	redisOnce sync.Once
)

// initRedis 初始化db连接
func initRedis() {
	redisConfig := configs.GetGlobalConfig().RedisConfig
	log.Infof("redisConfig:%+v", redisConfig)
	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}
	res, err := redisConn.Set(context.Background(), "abc", 100, 60).Result()
	log.Infof("res:%v,err:%v", res, err)
	_, err = redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to ping redis, err:%s")
	}
}

func CloseRedis() {
	if redisConn != nil {
		redisConn.Close()
	}
}

// GetRedisCli 获取数据库连接
func GetRedisCli() *redis.Client {
	redisOnce.Do(initRedis)

	return redisConn
}
