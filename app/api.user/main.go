/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/6
 */
package main

import (
	"context"
	"demo/app/api.user/conf"
	"demo/app/api.user/handler"
	"demo/app/api.user/service"
	"demo/cmn"
	"demo/utility/db"
	"demo/wrap"
	"fmt"
	"github.com/kataras/iris"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var serverAddress string

func main() {
	service.InitService(
		micro.Action(func(ctx *cli.Context) {
			// 初始化公共配置文件
			checkErr("InitCommonConfig", cmn.InitConfig(ctx.String("etcd_addr")))
			fmt.Print("InitCommonConfig Success!!!\n")

			// 初始化app配置文件
			checkErr("InitAppConfig", conf.InitConfig(ctx.String("etcd_addr"), cmn.APP_API_USER))
			fmt.Print("InitAppConfig Success!!!\n")

			// 获取接口的监听地址
			serverAddress = ctx.String("server_address")
		}),
	)

	// 创建文件日志，按天分割，日志文件仅保留一周
	w, err := rotatelogs.New(conf.GetLogPath())
	checkErr("CreateRotateLog", err)

	// 设置日志
	logrus.SetOutput(w)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("InitMysql...\r")
	checkErr("InitMysql", db.InitMysql(cmn.GetMysqlConfig()))
	fmt.Print("InitMysql Success!!!\n")

	// 启动redis
	defer db.CloseRedis()
	fmt.Print("InitRedis...\r")
	checkErr("InitRedis", db.InitRedis(cmn.GetRedisConfig()))
	fmt.Print("InitRedis Success!!!\n")

	app := iris.New()

	app.Use(wrap.RecoverMdwHandler())
	app.Use(wrap.AccessMdwHandler())

	// 优雅的关闭程序
	serverWG := new(sync.WaitGroup)
	defer serverWG.Wait()

	iris.RegisterOnInterrupt(func() {
		serverWG.Add(1)
		defer serverWG.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()

		_ = app.Shutdown(ctx)
	})

	// 注册路由
	app.Get("/ping", func(ctx iris.Context) { _, _ = ctx.WriteString("pong") })
	app.Get("/test", handler.ActionUsers)

	// server配置
	c := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           true,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: true,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
		IgnoreServerErrors:                []string{iris.ErrServerClosed.Error()},
		RemoteAddrHeaders:                 map[string]bool{"X-Real-Ip": true, "X-Forwarded-For": true},
	})

	_ = app.Run(iris.Addr(serverAddress), c)
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}
