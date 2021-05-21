package db

import (
	_ "gitee.com/itsos/golibs/tests"
	"golang.org/x/net/context"
	"testing"
)

func TestDds(t *testing.T) {
	Init()
	var ctx = context.Background()
	t.Log(Rdb.Set(ctx, "key", "value", 0).Err())
	t.Log(Rdb.Get(ctx, "key"))
}

func TestDd(t *testing.T) {
	Init()
	Conn.ShowExecTime(true)
}
