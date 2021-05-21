package core

import (
	"gitee.com/itsos/golibs/global/consts"
	"os"
)

func GetEnviron() string {
	env, ok := os.LookupEnv(consts.GoEnvironKey)
	if ok == false {
		env = "dev"
	}
	return env
}

func IsProductEnv() bool {
	return GetEnviron() == consts.EnvProduct
}

func GetTestBasePath() string {
	env, _ := os.LookupEnv(consts.TestBasePathKey)
	return env
}
