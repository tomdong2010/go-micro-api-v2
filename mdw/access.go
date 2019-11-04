package mdw

import (
	"context"
	"demo/proto"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"time"
)

// 记录请求日志
func LogMdw(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, _ := metadata.FromContext(ctx)

		begin := time.Now()

		defer func() {
			duration := time.Now().Sub(begin).Nanoseconds() / 1000000

			if r, ok := req.Body().(*proto.Request); ok {
				logrus.WithFields(logrus.Fields{
					"ip":       meta["Remote"],
					"method":   r.Method,
					"path":     r.Path,
					"header":   decodePair(r.Header),
					"queries":  decodePair(r.Get),
					"reqbody":  r.Body,
					"duration": duration,
				}).Info("123")
			} else {
				fmt.Printf("%#v\n", meta)
			}
		}()

		return fn(ctx, req, rsp)
	}
}

func decodePair(m map[string]*proto.Pair) map[string]string {
	var r = make(map[string]string, len(m))

	for k, v := range m {
		r[k] = v.GetValues()[0]
	}

	return r
}