package sqlite

import (
	"fmt"
	common2 "gitee.com/itsos/golibs/db/common"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct{}

func (s *sqlite) GetDsn() string {
	return fmt.Sprintf("%s?loc=%s", common2.Config.GetStorageFile(), common2.Config.GetTimezone())
}

const driver = "sqlite3"

func (s *sqlite) Connect() *common2.Dbs {
	common2.Config.UseSqlite()
	common2.Config.SetMode(common2.Master)
	engine, err := xorm.NewEngineGroup(driver, []string{s.GetDsn()})
	if err != nil {
		panic(err)
	}
	return &common2.Dbs{Conn: engine}
}

func NewSqlite() common2.Db {
	return &sqlite{}
}
