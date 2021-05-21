package mysql

import (
	"fmt"
	common2 "gitee.com/itsos/golibs/db/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// https://github.com/go-xorm/xorm
// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-01/2.engine_group.html#

type mysql struct{}

func (m *mysql) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		common2.Config.GetUser(),
		common2.Config.GetPassword(),
		common2.Config.GetHost(),
		common2.Config.GetPort(),
		common2.Config.GetDatabase(),
		common2.Config.GetCharset(),
	)
}

const driver = "mysql"

func (m *mysql) Connect() *common2.Dbs {
	var dataSourceNameSlice []string
	common2.Config.UseMysql()
	modes := []string{common2.Master, common2.Slave1, common2.Slave2, common2.Slave3}
	for _, mode := range modes {
		common2.Config.SetMode(mode)
		if common2.Config.GetHost() != "" {
			dataSourceNameSlice = append(dataSourceNameSlice, m.GetDsn())
		}
	}
	engine, err := xorm.NewEngineGroup(driver, dataSourceNameSlice)
	engine.TZLocation, _ = time.LoadLocation(common2.Config.GetTimezone())
	engine.DatabaseTZ, _ = time.LoadLocation(common2.Config.GetTimezone())
	if err != nil {
		panic(err)
	}
	return &common2.Dbs{Conn: engine}
}

func NewMysql() common2.Db {
	return &mysql{}
}
