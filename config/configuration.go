/*
   Copyright (c) [2021] itsos
   golibs is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package config

import (
	"fmt"
	"gitee.com/itsos/golibs/v2/global/variable"
	"gitee.com/itsos/golibs/v2/utils/reflects"
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
	GetTimezone() string
	GetSock() string
	GetSwaggerPort() string
	GetScheme() string
	GetEs() []string
	GetMysql() map[string]IMysql
	GetSqlite() ISqlite
	GetRedis() IRedis
}

type Configuration struct {
	Active      string `yaml:"active"`
	Domain      string `yaml:"domain"`
	Port        string `yaml:"port"`
	Scheme      string `yaml:"scheme"`
	Sock        string `yaml:"sock"`
	Timezone    string `yaml:"timezone"`
	SwaggerPort string `yaml:"swagger.port"`
	Es          string `yaml:"es"`
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

// GetEs 获取es配置
func (c Configuration) GetEs() []string {
	return viper.GetStringSlice(c.Es)
}

func (c Configuration) GetSock() string {
	return viper.GetString(c.Sock)
}

func (c Configuration) GetTimezone() string {
	return viper.GetString(c.Timezone)
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

// sqliteAlone sqlite配置实例
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

// mysqlAlone mysql配置实例
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

// redisAlone redis配置实例
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

// mysql mysql配置
type mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
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

type IMysql interface {
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabase() string
	GetCharset() string
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

// CovertConfiguration struct tag 指定为属性值，方便操作
func CovertConfiguration() *Configuration {
	loadConfigFile()
	c := &Configuration{}
	t := reflect.ValueOf(c).Elem()
	reflects.TagToValueFlip(t, reflects.YAML)
	return c
}

// loadConfigFile 加载配置文件
func loadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(variable.ConfPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

var Config = CovertConfiguration()

// minio
type minio struct {
	Endpoint        string `yaml:"minio.endpoint"`
	AccessKeyID     string `yaml:"minio.access_key_id"`
	SecretAccessKey string `yaml:"minio.secret_access_key"`
	UseSSL          string `yaml:"minio.use_ssl"`
}

func (m minio) GetEndpoint() string {
	return viper.GetString(m.Endpoint)
}

func (m minio) GetAccessKeyID() string {
	return viper.GetString(m.AccessKeyID)
}

func (m minio) GetSecretAccessKey() string {
	return viper.GetString(m.SecretAccessKey)
}

func (m minio) GetUseSSL() bool {
	return viper.GetBool(m.UseSSL)
}

type IMinio interface {
	GetEndpoint() string
	GetAccessKeyID() string
	GetSecretAccessKey() string
	GetUseSSL() bool
}

var minioIns IMinio = nil

func GetMinio() IMinio {
	if minioIns == nil {
		minioIns = &minio{}
		reflects.TagToValueFlip(reflect.ValueOf(minioIns).Elem(), reflects.YAML)
	}
	return minioIns
}
