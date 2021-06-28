package redis

import (
	"fmt"
	"gitee.com/itsos/golibs/db/common"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis

type rds struct{}

func (r *rds) GetDsn() string {
	return fmt.Sprintf("%s:%d", common.Config.GetHost(), common.Config.GetPort())
}

var ctx = context.Background()

func (r *rds) Connect() *common.Dbs {
	common.Config.UseRedis()
	common.Config.SetMode(common.Master)

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.GetDsn(),
		Password: common.Config.GetPassword(),
		DB:       common.Config.GetDb(),
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	return &common.Dbs{Rdb: rdb}
}

func NewRedis() common.Db {
	return &rds{}
}
