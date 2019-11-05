package wrap

import (
	"context"
	"demo/utility/helper"
	"github.com/google/uuid"
	"github.com/kataras/iris"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"time"
)

// micro微服务记录请求日志
func AccessWrapHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		begin := time.Now()

		meta, _ := metadata.FromContext(ctx)

		defer func() {
			logrus.WithFields(logrus.Fields{
				"ip":       "",
				"method":   "RPC",
				"path":     meta["Micro-Service"] + "." + meta["Micro-Method"],
				"reqId":    meta["Request-Id"],
				"header":   meta,
				"queries":  "",
				"reqBody":  req.Body(),
				"duration": time.Now().Sub(begin).Milliseconds(),
			}).Info(rsp)
		}()

		return fn(ctx, req, rsp)
	}
}

// iris接口记录请求日志
func AccessMdwHandler() iris.Handler {
	return func(ctx iris.Context) {
		begin := time.Now()

		if reqId := ctx.GetHeader("Request-Id"); reqId == "" {
			ctx.Values().Set("Request-Id", uuid.New())
		} else {
			ctx.Values().Set("Request-Id", reqId)
		}

		ctx.Values().Set("lang", ctx.URLParamDefault("lang", "en"))

		ctx.Record()

		defer func() {
			logrus.WithFields(logrus.Fields{
				"ip":       ctx.RemoteAddr(),
				"method":   ctx.Method(),
				"path":     ctx.Path(),
				"reqId":    ctx.Values().GetString("Request-Id"),
				"header":   helper.RequestHeader(ctx),
				"queries":  helper.RequestQueries(ctx),
				"reqBody":  helper.RequestBody(ctx),
				"duration": time.Now().Sub(begin).Milliseconds(),
			}).Info(helper.RequestBody(ctx))
		}()

		ctx.Next()
	}
}
