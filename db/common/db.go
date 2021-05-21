package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-xorm/xorm"
)

type Db interface {
	GetDsn() string
	Connect() *Dbs
}

type Dbs struct {
	Conn *xorm.EngineGroup
	Rdb  *redis.Client
}
