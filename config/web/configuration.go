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
	"github.com/spf13/viper"
	"reflect"
)

// web 应用基础配置

type ConfigurationReadOnly interface {
	GetActive() string
	GetUrl() string
	GetSwaggerUrl() string
	GetDomain() string
	GetPort() string
	GetSwaggerPort() string
	GetScheme() string
	GetEs() []string
	GetSock() string
}

type Configuration struct {
	Active      string `yaml:"active"`
	Domain      string `yaml:"domain"`
	Port        string `yaml:"port"`
	Scheme      string `yaml:"scheme"`
	SwaggerPort string `yaml:"swagger.port"`
	Es          string `yaml:"es"`
	Sock        string `yaml:"sock"`
}

func (c Configuration) GetUrl() string {
	url := c.GetScheme() + "://" + c.GetDomain()
	if c.GetPort() != "" {
		url += ":" + c.GetPort()
	}
	return url
}

func (c Configuration) GetSwaggerUrl() string {
	url := c.GetScheme() + "://" + c.GetDomain()
	if c.GetSwaggerPort() != "" {
		url += ":" + c.GetSwaggerPort()
	}
	return url
}

func (c Configuration) GetEs() []string {
	return viper.GetStringSlice(c.Es)
}

func (c Configuration) GetSock() string {
	return viper.GetString(c.Sock)
}

func (c Configuration) GetActive() string {
	return viper.GetString(c.Active)
}

func (c Configuration) GetDomain() string {
	return viper.GetString(c.Domain)
}

func (c Configuration) GetPort() string {
	return viper.GetString(c.Port)
}

func (c Configuration) GetSwaggerPort() string {
	return viper.GetString(c.SwaggerPort)
}

func (c Configuration) GetScheme() string {
	return viper.GetString(c.Scheme)
}

var _ ConfigurationReadOnly = (*Configuration)(nil)

func CovertConfiguration() *Configuration {
	c := &Configuration{}
	t := reflect.ValueOf(c).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Type().Field(i)
		tag := f.Tag.Get("yaml")
		s := t.FieldByName(f.Name)
		if !s.CanSet() {
			panic(f.Name + " => not set value. " + tag)
		}
		s.SetString(tag)
	}
	return c
}

var Config = CovertConfiguration()
