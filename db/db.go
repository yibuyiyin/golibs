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

type GoLibDb = *xorm.EngineGroup

var Conn *xorm.EngineGroup
var Rdb *redis2.Client
var once sync.Once

// Init 兼容老的方法，已过时
func Init() {
	once.Do(func() {
		if common.Config.Init().GetStorageDriver() == common.DriverMysql {
			Conn = mysql.NewMysqlOld().Connect().Conn
		} else {
			Conn = sqlite.NewSqliteOld().Connect().Conn
		}
		Rdb = redis.NewRedisOld().Connect().Rdb
	})
}
