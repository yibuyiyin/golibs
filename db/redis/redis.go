package redis

import (
	"fmt"
	"gitee.com/itsos/golibs/v2/config"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"sync"
)

// https://github.com/go-redis/redis

type GoLibRedis = *redis.Client

var redisOnce sync.Once
var redisNew GoLibRedis

func NewRedis() GoLibRedis {
	redisOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%d", config.Config.GetRedis().GetHost(), config.Config.GetRedis().GetPort())
		redisNew = redis.NewClient(&redis.Options{
			Addr:     dsn,
			Password: config.Config.GetRedis().GetPassword(),
			DB:       config.Config.GetRedis().GetDb(),
		})
		err := redisNew.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	})
	return redisNew
}
