/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/6
 */
package main

import (
	"demo/app/api.gateway/conf"
	"demo/app/api.gateway/handler"
	"demo/app/api.gateway/service"
	common "demo/conf"
	"demo/utility/db"
	"demo/utility/helper"
	"demo/utility/log"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/valyala/fasthttp"
)

var serverAddress string
var appName = common.APP_API_GATEWAY

func main() {
	defer uninit()
	service.InitService(appName, micro.Action(initialize))

	h := fasthttp.CompressHandler(handler.Handler)
	helper.CheckErr("FastHttp Start", fasthttp.ListenAndServe(serverAddress, h), true)
}

func initialize(ctx *cli.Context) {
	// 初始化公共配置文件
	helper.CheckErr("InitCommonConfig", common.InitConfig(ctx.String("etcd_addr")), true)

	// 初始化app配置文件
	helper.CheckErr("InitAppConfig", conf.InitConfig(ctx.String("etcd_addr"), appName), true)

	// 获取接口的监听地址
	if serverAddress = ctx.String("server_address"); serverAddress == "" {
		serverAddress = "0.0.0.0:8080"
	}

	// 初始化日志
	helper.CheckErr("InitZapLog", log.InitZapLogger(conf.GetLogPath()), true)

	// 启动mysql
	helper.CheckErr("InitMysql", db.InitMysql(common.GetMysqlConfig()), true)

	// 启动redis
	helper.CheckErr("InitRedis", db.InitRedis(common.GetRedisConfig()), true)
}

func uninit() {
	db.CloseMysql()
	db.CloseRedis()
}
