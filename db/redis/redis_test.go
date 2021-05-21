package redis

import (
	_ "gitee.com/itsos/golibs/tests"
	"github.com/go-redis/redis/v8"
	"testing"
)

func initConfig() *redis.Client {
	db := NewRedis().Connect().Rdb
	return db
}

func TestInitConfig(t *testing.T) {
	initConfig()
}
