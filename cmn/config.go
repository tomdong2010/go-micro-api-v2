package cmn

import (
	. "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
	"strings"
	"time"
)

var conf *config

// 公共配置
type config struct {
	MicroConf Config
}

// 初始化公共配置
func InitConfig(etcd_addr string) error {
	var err error

	conf = &config{NewConfig()}

	if etcd_addr != "" {
		err = conf.MicroConf.Load(etcd.NewSource(
			etcd.WithAddress(strings.Split(etcd_addr, ",")...),
			etcd.WithPrefix(APP_CONF_PREFIX+"/common"),
			etcd.StripPrefix(true),
			etcd.WithDialTimeout(10 * time.Second),
		))
	} else {
		err = conf.MicroConf.Load(file.NewSource(file.WithPath("common.yaml")))
	}

	return err
}

// 获取环境变量
func GetEnv() string {
	return conf.MicroConf.Get("env").String(ENV_DEV)
}

// 获取mysql配置
func GetMysqlConfig() (string, int, int) {
	dsn := conf.MicroConf.Get("mysql", "dsn").String("")
	maxIdle := conf.MicroConf.Get("mysql", "max_idle").Int(1)
	maxOpen := conf.MicroConf.Get("mysql", "max_open").Int(10)

	return dsn, maxIdle, maxOpen
}

// 获取redis配置
func GetRedisConfig() (string, int, int, int) {
	addr := conf.MicroConf.Get("redis", "addr").String("")
	db := conf.MicroConf.Get("redis", "db").Int(0)
	maxIdle := conf.MicroConf.Get("redis", "max_idle").Int(1)
	maxOpen := conf.MicroConf.Get("redis", "max_open").Int(10)

	return addr, db, maxIdle, maxOpen
}
