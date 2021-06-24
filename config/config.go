package config

import (
	"fmt"
	"gitee.com/itsos/golibs/global/variable"
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

func Init() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(variable.BasePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
}
