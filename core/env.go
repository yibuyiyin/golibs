package core

import (
	"gitee.com/itsos/golibs/v2/global/consts"
	"os"
)

func GetDevBasePath() string {
	env, _ := os.LookupEnv(consts.DevBasePathKey)
	return env
}
