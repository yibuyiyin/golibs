package variable

import (
	"flag"
	"gitee.com/itsos/golibs/v2/core"
	"gitee.com/itsos/golibs/v2/global/consts"
	"os"
	"strings"
)

var (
	// BasePath 开发项目根目录
	BasePath = ""
)

var workdir = flag.String("w", "", "指定工作目录 -w /workdir")

func init() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
		BasePath = core.GetDevBasePath()
		if BasePath == "" {
			panic("Environment variable [" + consts.DevBasePathKey + "] not set. Specify the project root directory.")
		}
	} else {
		flag.Parse()
		if *workdir != "" {
			BasePath = *workdir
		} else if pwd, err := os.Getwd(); err == nil {
			BasePath = pwd
		}
	}
}
