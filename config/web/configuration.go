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
	"gitee.com/itsos/golibs/config"
	"github.com/goinggo/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
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
	GetMysql() map[string]IMysql
	GetSqlite() ISqlite
	GetRedis() IRedis
}

type Configuration struct {
	Active      string `yaml:"active"`
	Domain      string `yaml:"domain"`
	Port        string `yaml:"port"`
	Scheme      string `yaml:"scheme"`
	SwaggerPort string `yaml:"swagger.port"`
	Es          string `yaml:"es"`
	Sock        string `yaml:"sock"`
	Mysql       string `yaml:"mysql"`
	Sqlite      string `yaml:"sqlite"`
	Redis       string `yaml:"redis"`
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

var sqliteAlone ISqlite

func (c Configuration) GetSqlite() ISqlite {
	if sqliteAlone == nil {
		v := viper.GetStringMapString(c.Sqlite)
		s := sqlite{}
		if err := mapstructure.Decode(v, &s); err != nil {
			panic(err)
		}
		sqliteAlone = s
	}
	return sqliteAlone
}

var redisAlone IRedis

func (c Configuration) GetRedis() IRedis {
	if redisAlone == nil {
		v := viper.GetStringMapString(c.Redis)
		r := redis{}
		if err := mapstructure.Decode(v, &r); err != nil {
			panic(err)
		}
		redisAlone = r
	}
	return redisAlone
}

var mysqlAlone map[string]IMysql

func (c Configuration) GetMysql() map[string]IMysql {
	if mysqlAlone == nil {
		mysqlAlone = make(map[string]IMysql)
		for k, v := range viper.GetStringMap(c.Mysql) {
			m := mysql{}
			if err := mapstructure.Decode(v, &m); err != nil {
				panic(err)
			}
			mysqlAlone[k] = m
		}
	}
	return mysqlAlone
}

// mysql mysql配置
type mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
	Timezone string
}

func (m mysql) GetHost() string {
	return m.Host
}

func (m mysql) GetPort() int {
	return m.Port
}

func (m mysql) GetUser() string {
	return m.User
}

func (m mysql) GetPassword() string {
	return m.Password
}

func (m mysql) GetDatabase() string {
	return m.Database
}

func (m mysql) GetCharset() string {
	return m.Charset
}

func (m mysql) GetTimezone() string {
	return m.Timezone
}

type IMysql interface {
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabase() string
	GetCharset() string
	GetTimezone() string
}

var _ IMysql = (*mysql)(nil)

// sqlite sqlite配置
type sqlite struct {
	Timezone    string
	StorageFile string
}

func (s sqlite) GetTimezone() string {
	return s.Timezone
}

func (s sqlite) GetStorageFile() string {
	return s.StorageFile
}

type ISqlite interface {
	GetTimezone() string
	GetStorageFile() string
}

var _ ISqlite = (*sqlite)(nil)

// redis redis配置
type redis struct {
	Host     string
	Port     string
	Password string
	Db       string
}

func (r redis) GetHost() string {
	return r.Host
}

func (r redis) GetPort() int {
	port, _ := strconv.Atoi(r.Port)
	return port
}

func (r redis) GetPassword() string {
	return r.Password
}

func (r redis) GetDb() int {
	db, _ := strconv.Atoi(r.Db)
	return db
}

type IRedis interface {
	GetHost() string
	GetPort() int
	GetPassword() string
	GetDb() int
}

var _ IRedis = (*redis)(nil)

var _ ConfigurationReadOnly = (*Configuration)(nil)

func CovertConfiguration() *Configuration {
	config.Init()
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
