package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	global "nft/server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(1)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,                 // use default DB
		PoolSize: redisConf.PoolSize, //连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error("redis连接失败", redisConf.Addr())

		return nil
	}
	return rdb
}
