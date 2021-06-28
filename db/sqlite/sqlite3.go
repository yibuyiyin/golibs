package sqlite

import (
	"fmt"
	"gitee.com/itsos/golibs/config/web"
	"gitee.com/itsos/golibs/db/common"
	"gitee.com/itsos/golibs/global/consts"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
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

func NewSqlite() common.Db {
	return &sqlite{}
}
