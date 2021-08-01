package sqlite

import (
	"fmt"
	"gitee.com/itsos/golibs/config/web"
	"gitee.com/itsos/golibs/db/common"
	"gitee.com/itsos/golibs/global/consts"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type sqlite struct{}

func (s *sqlite) GetDsn() string {
	return fmt.Sprintf("%s?loc=%s", common.Config.GetStorageFile(), common.Config.GetTimezone())
}

const driver = "sqlite3"

func (s *sqlite) Connect() *common.Dbs {
	common.Config.UseSqlite()
	common.Config.SetMode(common.Master)
	engine, err := xorm.NewEngineGroup(driver, []string{s.GetDsn()})
	if web.Config.GetActive() == consts.EnvDev {
		engine.ShowSQL(true)
	}
	if err != nil {
		panic(err)
	}
	return &common.Dbs{Conn: engine}
}

func NewSqliteOld() common.Db {
	return &sqlite{}
}

var sqliteOnce sync.Once
var sqliteNew *xorm.EngineGroup

func NewSqlite() *xorm.EngineGroup {
	sqliteOnce.Do(func() {
		sqliteNew = NewSqliteOld().Connect().Conn
	})
	return sqliteNew
}
