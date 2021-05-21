package db

import (
	"gitee.com/itsos/golibs/db/common"
	"gitee.com/itsos/golibs/db/mysql"
	"gitee.com/itsos/golibs/db/redis"
	"gitee.com/itsos/golibs/db/sqlite"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
	"sync"
)

var Conn *xorm.EngineGroup
var Rdb *redis2.Client
var once sync.Once

func Init() {
	once.Do(func() {
		if common.Config.Init().GetStorageDriver() == common.DriverMysql {
			Conn = mysql.NewMysql().Connect().Conn
		} else {
			Conn = sqlite.NewSqlite().Connect().Conn
		}
		Rdb = redis.NewRedis().Connect().Rdb
	})
}
