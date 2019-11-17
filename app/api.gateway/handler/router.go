package handler

import (
	admin_user "demo/app/api.gateway/handler/admin.user"
	"demo/utility/helper"
	"demo/utility/log"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"runtime"
	"strings"
	"time"
)

// 设置路由
func routes() {
	router.GET("/admin/user/login", admin_user.UserLogin)
}

var router *fasthttprouter.Router

func init() {
	router = fasthttprouter.New()
	router.PanicHandler = recoverHandler

	routes()
}

func recoverHandler(ctx *fasthttp.RequestCtx, v interface{}) {
	var trace string
	for i := 1; ; i++ {
		if _, f, l, got := runtime.Caller(i); !got {
			break
		} else if strings.Contains(f, "/app/") {
			trace += fmt.Sprintf("%s:%d;", f, l)
		}
	}

	log.Error("recover exception.",
		zap.ByteString("method", ctx.Method()),
		zap.ByteString("path", ctx.Path()),
		zap.Error(v.(error)),
		zap.String("trace", trace))
}

func Handler(ctx *fasthttp.RequestCtx) {
	begin := time.Now()

	if reqId := ctx.Request.Header.Peek("Request-Id"); len(reqId) == 0 {
		ctx.SetUserValue("Request-Id", uuid.New().String())
	} else {
		ctx.SetUserValue("Request-Id", string(reqId))
	}

	defer func() {
		log.Info("access log.",
			zap.ByteString("ip", helper.RealIpAddr(ctx)),
			zap.ByteString("method", ctx.Method()),
			zap.ByteString("path", ctx.Path()),
			zap.String("reqId", ctx.UserValue("Request-Id").(string)),
			zap.String("queries", ctx.QueryArgs().String()),
			zap.ByteString("reqBody", ctx.PostBody()),
			zap.ByteString("respBody", ctx.Response.Body()),
			zap.Duration("duration", time.Now().Sub(begin)))
	}()

	router.Handler(ctx)
}
