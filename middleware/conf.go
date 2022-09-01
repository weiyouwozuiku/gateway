package middleware

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var ConfEnv string
var DBMapPool map[string]*sql.DB
var GORMMapPool map[string]*gorm.DB

type MysqlMapConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

func GetConfEnv() string {
	return ConfEnv
}

func ParseConfig(path string, conf any) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open config %v fail,%v", path, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read config fail,%v", err)
	}
	v := viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(data))
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("Parse config fail,config:%v, err:%v", string(data), err)
	}
	return nil
}
