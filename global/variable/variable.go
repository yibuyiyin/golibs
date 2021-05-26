package variable

import (
	"gitee.com/itsos/golibs/core"
	"gitee.com/itsos/golibs/global/consts"
	"os"
	"strings"
)

var (
	// BasePath 开发项目根目录
	BasePath = core.GetDevBasePath()
)

func init() {
	if dir, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			if BasePath == "" {
				panic("Environment variable [" + consts.DevBasePathKey + "] not set")
			}
		} else {
			BasePath = dir
		}
	}
}
