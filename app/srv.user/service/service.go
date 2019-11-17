package service

import (
	"context"
	"demo/utility/log"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"runtime"
	"strings"
	"time"
)

var microService micro.Service

// 初始化service
func InitService(appName string, opts... micro.Option) {
	microService = micro.NewService(
		micro.Name(appName),
		micro.WrapHandler(recoverHandler),
		micro.WrapHandler(accessHandler),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.Flags(cli.StringFlag{
			Name:   "etcd_addr",
			EnvVar: "ETCD_ADDR",
			Usage:  "This is etcd config address.",
		}),
	)

	microService.Init(opts...)
}

func Run() error {
	return microService.Run()
}

func Server() server.Server {
	return microService.Server()
}

// 创建上下文
func NewContext(ctx *fasthttp.RequestCtx) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{
		"Request-Id": ctx.UserValue("Request-Id").(string),
	})
}

func NewClient() client.Client {
	return microService.Client()
}

func recoverHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		defer func() {
			if err := recover(); err != nil {
				var trace string
				for i := 1; ; i++ {
					if _, f, l, got := runtime.Caller(i); !got {
						break
					} else if strings.Contains(f, "/app/") {
						trace += fmt.Sprintf("%s:%d;", f, l)
					}
				}

				log.Error("recover exception.",
					zap.String("server", req.Endpoint()),
					zap.Any("reqBody", req.Body()),
					zap.Any("error", err),
					zap.String("trace", trace))
			}
		}()
		return fn(ctx, req, rsp)
	}
}

func accessHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		begin := time.Now()

		meta, _ := metadata.FromContext(ctx)

		defer func() {
			log.Info("access log.",
				zap.String("ip", meta["Ip-Addr"]),
				zap.String("method", "RPC"),
				zap.String("path", meta["Micro-Service"] + "." + meta["Micro-Method"]),
				zap.String("reqId", meta["Request-Id"]),
				zap.String("queries", ""),
				zap.Any("reqBody", req.Body()),
				zap.Any("respBody", rsp),
				zap.Duration("duration", time.Now().Sub(begin)))
		}()
		return fn(ctx, req, rsp)
	}
}