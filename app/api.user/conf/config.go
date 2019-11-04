/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/8/30
 */
package conf

import (
	"demo/cmn"
	"errors"
	. "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/config/source/consul"
)

var conf *config

// 项目配置
type config struct {
	AppName string
	MicroConf Config
}

// 初始化项目配置
func InitConfig(service server.Server, appName string) error {
	var err error

	conf = &config{appName, NewConfig()}

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
			etcd.WithPrefix(cmn.APP_CONF_PREFIX + "/" + appName),
			etcd.StripPrefix(true),
		))
	case "consul":
		err = conf.MicroConf.Load(consul.NewSource(
			consul.WithAddress(registryAddr[0]),
			consul.WithPrefix(cmn.APP_CONF_PREFIX + "/" + appName),
			consul.StripPrefix(true),
		))
	default:
		err = conf.MicroConf.Load(file.NewSource(file.WithPath(appName + ".yaml")))
	}

	return err
}

func GetLogPath() string {
	return conf.MicroConf.Get("log_path").String(conf.AppName + ".%Y%m%d.log")
}

func GetUserPrefix() string {
	return conf.MicroConf.Get("user_prefix").String("default_prefix")
}
