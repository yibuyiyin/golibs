package variable

import (
	core2 "gitee.com/itsos/golibs/core"
	consts2 "gitee.com/itsos/golibs/global/consts"
	"os"
	"strings"
)

var (
	// BasePath 项目根目录
	BasePath = core2.GetTestBasePath()
)

func init() {
	if dir, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			if BasePath == "" {
				panic("Environment variable [" + consts2.TestBasePathKey + "] not set")
			}
		} else {
			BasePath = dir
		}
	}
}
