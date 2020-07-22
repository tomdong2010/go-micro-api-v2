/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/8/30
 */
package conf

import (
	. "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
	"strings"
	"time"
)

var conf *config

// 项目配置
type config struct {
	AppName   string
	MicroConf Config
}

// 初始化项目配置
func InitConfig(etcd_addr, appName string) error {
	var err error
	con, _ := NewConfig()
	conf = &config{appName, con}

	if etcd_addr != "" {
		err = conf.MicroConf.Load(etcd.NewSource(
			etcd.WithAddress(strings.Split(etcd_addr, ",")...),
			etcd.WithPrefix(appName),
			etcd.StripPrefix(true),
			etcd.WithDialTimeout(10*time.Second),
		))
	} else {
		err = conf.MicroConf.Load(file.NewSource(file.WithPath(appName + ".yaml")))
	}

	return err
}

func GetLogPath() string {
	return conf.MicroConf.Get("log", "path").String(conf.AppName + ".%Y%m%d.log")
}
