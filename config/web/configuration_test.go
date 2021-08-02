/*
   Copyright (c) [2021] itsos
   golibs is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package web

import (
	"testing"
)

func TestCovertConfiguration(t *testing.T) {
	t.Log(Config.GetUrl())
	t.Log(Config.GetSwaggerUrl())
	t.Log(Config.GetActive())
	t.Log(Config.GetEs())
}

func TestGetMysql(t *testing.T) {
	t.Log(Config.GetMysql())
	t.Log(Config.GetMysql()["slave1"].GetHost())
}

func TestGetSqlite(t *testing.T) {
	t.Log(Config.GetSqlite())
	t.Log(Config.GetSqlite().GetStorageFile())
}

func TestGetRedis(t *testing.T) {
	t.Log(Config.GetRedis().GetHost())
	t.Log(Config.GetRedis().GetPort())
	t.Log(Config.GetRedis().GetPassword())
	t.Log(Config.GetRedis().GetDb())
}

func BenchmarkConfiguration_GetMysql(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Config.GetMysql()
			Config.GetMysql()["slave1"].GetHost()
		}
	})
}
