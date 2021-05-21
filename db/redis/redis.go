package redis

import (
	"fmt"
	common2 "gitee.com/itsos/golibs/db/common"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

// https://github.com/go-redis/redis

type rds struct{}

func (r *rds) GetDsn() string {
	return fmt.Sprintf("%s:%d", common2.Config.GetHost(), common2.Config.GetPort())
}

var ctx = context.Background()

func (r *rds) Connect() *common2.Dbs {
	common2.Config.UseRedis()
	common2.Config.SetMode(common2.Master)

	rdb := redis.NewClient(&redis.Options{
		Addr:     r.GetDsn(),
		Password: common2.Config.GetPassword(),
		DB:       common2.Config.GetDb(),
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	return &common2.Dbs{Rdb: rdb}
}

func NewRedis() common2.Db {
	return &rds{}
}
