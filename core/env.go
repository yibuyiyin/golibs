package core

import (
	"gitee.com/itsos/golibs/global/consts"
	"os"
)

func GetDevBasePath() string {
	env, _ := os.LookupEnv(consts.DevBasePathKey)
	return env
}
