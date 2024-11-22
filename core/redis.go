package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	global "nft/server/gloabl"
	"time"
)

func ConnectRedis() *redis.Client {
<<<<<<< HEAD
<<<<<<< HEAD
	return ConnectRedisDB(1)
=======
	return ConnectRedisDB(0)
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
=======
	return ConnectRedisDB(0)
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
<<<<<<< HEAD
<<<<<<< HEAD
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,                 // use default DB
=======
		Addr:     "192.168.189.11:6379",
		Password: "",
		DB:       0,                  // use default DB
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
=======
		Addr:     "192.168.189.11:6379",
		Password: "",
		DB:       0,                  // use default DB
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
		PoolSize: redisConf.PoolSize, //连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		logrus.Error("redis连接失败", redisConf.Addr())
=======
		logrus.Error("redis连接失败 %s", redisConf.Addr())
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
=======
		logrus.Error("redis连接失败 %s", redisConf.Addr())
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
		return nil
	}
	return rdb
}
