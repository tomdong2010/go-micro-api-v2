package service

import (
	"context"
	"demo/cmn"
	"github.com/kataras/iris"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

var microService micro.Service

// 初始化service
func InitService(opts... micro.Option) {
	microService = micro.NewService(
		micro.Name(cmn.APP_NAME_PREFIX+cmn.APP_API_USER),
		micro.Flags(cli.StringFlag{
			Name:   "etcd_addr",
			EnvVar: "ETCD_ADDR",
			Usage:  "This is etcd config address.",
		}),
	)

	microService.Init(opts...)
}

// 创建上下文
func NewContext(ctx iris.Context) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		"Request-Id": ctx.Values().GetString("Request-Id"),
	})
}
