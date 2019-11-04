package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"demo/utility/helper"
	"time"
)

var mysqlDB *gorm.DB

// 初始化mysql
func InitMysql(conf map[string]string) (err error){
	if !helper.MapSKeysExists(conf, []string{"dsn", "max_idle", "max_open"}) {
		return errors.New("mysql config is missing")
	}

	mysqlDB, err = gorm.Open("mysql", conf["dsn"])

	if err == nil {
		mysqlDB.DB().SetMaxIdleConns(helper.StrToInt(conf["max_idle"], 1))
		mysqlDB.DB().SetMaxOpenConns(helper.StrToInt(conf["max_open"], 10))
		mysqlDB.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	}

	return
}

// 获取mysql连接
func GetMysql() *gorm.DB {
	return mysqlDB
}

// 关闭mysql
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}
