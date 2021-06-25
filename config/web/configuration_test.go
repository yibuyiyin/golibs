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
	_ "gitee.com/itsos/golibs/tests"
	"testing"
)

func TestCovertConfiguration(t *testing.T) {
	c := CovertConfiguration()
	t.Log(c.GetUrl())
	t.Log(c.GetSwaggerUrl())
	t.Log(c.GetActive())
}