package cmn

import (
	"errors"
	. "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/config/source/consul"
)

var conf *config

// 公共配置
type config struct {
	MicroConf Config
}

// 初始化公共配置
func InitConfig(service server.Server) error {
	var err error

	conf = &config{NewConfig()}

	registryName := service.Options().Registry.String()
	registryAddr := service.Options().Registry.Options().Addrs

	// 判断注册地址是否为空
	if registryName != "mdns" && len(registryAddr) == 0 {
		return errors.New("config path is required.")
	}

	// 目前支持etcd、consul和file三种
	switch registryName {
	case "etcd":
		err = conf.MicroConf.Load(etcd.NewSource(
			etcd.WithAddress(registryAddr...),
			etcd.WithPrefix(APP_CONF_PREFIX+"/common"),
			etcd.StripPrefix(true),
		))
	case "consul":
		err = conf.MicroConf.Load(consul.NewSource(
			consul.WithAddress(registryAddr[0]),
			consul.WithPrefix(APP_CONF_PREFIX+"/common"),
			consul.StripPrefix(true),
		))
	default:
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
