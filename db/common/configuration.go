package common

import "github.com/spf13/viper"

type Configuration interface {
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabase() string
	GetDb() int
	GetCharset() string
	GetTimezone() string
	GetStorageFile() string
	GetStorageDriver() string
	SetMode(mode string)
	Init() Configuration
	UseRedis()
	UseMysql()
	UseSqlite()
	parse()
}

type configuration struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	Db            int    `yaml:"db"`
	Database      string `yaml:"database"`
	Charset       string `yaml:"charset"`
	Timezone      string `yaml:"timezone"`
	StorageFile   string `yaml:"storage-file"`
	StorageDriver string `yaml:"storage-driver"`
	mode          string
	storageType   string
}

const (
	host          = "host"
	port          = "port"
	user          = "user"
	password      = "password"
	database      = "database"
	db            = "db"
	charset       = "charset"
	storageFile   = "storage-file"
	storageDriver = "storage-driver"
	timezone      = "timezone"
)

func (c *configuration) parse() {
	c.Host = c.getFieldString(host)
	c.Port = c.getFieldInt(port)
	c.User = c.getFieldString(user)
	c.Password = c.getFieldString(password)
	c.Database = c.getFieldString(database)
	c.Db = c.getFieldInt(db)
	c.Charset = c.getFieldString(charset)
	c.StorageFile = c.getFieldString(storageFile)
	c.Timezone = c.getFieldString(timezone)
}

func (c *configuration) getFieldString(field string) string {
	return viper.GetString(c.getStorageType() + "." + c.getMode() + "." + field)
}

func (c *configuration) getRootFieldString(field string) string {
	return viper.GetString(field)
}

func (c *configuration) getFieldInt(field string) int {
	return viper.GetInt(c.getStorageType() + "." + c.getMode() + "." + field)
}

const (
	Master = "master"
	Slave1 = "slave1"
	Slave2 = "slave2"
	Slave3 = "slave3"
)

// SetMode 设置mode 见上
func (c *configuration) SetMode(mode string) {
	c.mode = mode
	c.parse()
}

const (
	DriverRedis  = "redis"
	DriverSqlite = "sqlite3"
	DriverMysql  = "mysql"
)

func (c *configuration) UseRedis() {
	c.storageType = DriverRedis
}

func (c *configuration) UseSqlite() {
	c.storageType = DriverSqlite
}

func (c *configuration) UseMysql() {
	c.storageType = DriverMysql
}

func (c *configuration) getMode() string {
	return c.mode
}

func (c *configuration) getStorageType() string {
	return c.storageType
}

func (c *configuration) getStorageFile() string {
	return c.StorageFile
}

func (c *configuration) getTimezone() string {
	return c.Timezone
}

func (c *configuration) GetHost() string {
	return c.Host
}

func (c *configuration) GetPort() int {
	return c.Port
}

func (c *configuration) GetPassword() string {
	return c.Password
}

func (c *configuration) GetDatabase() string {
	return c.Database
}

func (c *configuration) GetDb() int {
	return c.Db
}

func (c *configuration) GetCharset() string {
	return c.Charset
}

func (c *configuration) GetUser() string {
	return c.User
}

func (c *configuration) GetStorageDriver() string {
	return c.StorageDriver
}

func (c *configuration) Init() Configuration {
	c.StorageDriver = c.getRootFieldString(storageDriver)
	return c
}

func (c *configuration) GetTimezone() string {
	return c.Timezone
}

func (c *configuration) GetStorageFile() string {
	return c.StorageFile
}

func newConfiguration() Configuration {
	return &configuration{}
}

var Config = newConfiguration()
