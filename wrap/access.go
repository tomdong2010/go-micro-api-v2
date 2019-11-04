package wrap

import (
	"context"
	"demo/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"time"
)

// 记录请求日志
func AccessWrapHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		var (
			reqId    string
			ipAddr   string
			duration int64
			method   string
			path     string
			header   interface{}
			queries  interface{}
			body     interface{}
			result   interface{}
		)

		begin := time.Now()

		meta, _ := metadata.FromContext(ctx)

		if v, ok := meta["Remote"]; ok {
			ipAddr = v
		} else {
			ipAddr = "0.0.0.0"
		}

		if v, ok := meta["Request-Id"]; ok {
			reqId = v
		} else {
			reqId = uuid.New().String()
		}

		if r, ok := req.Body().(*proto.Request); ok {
			method = r.Method
			path = r.Path
			queries = decodePair(r.Get)
			header = decodePair(r.Header)

			if v, ok := r.Header["Content-Type"]; ok && v.Values[0] == "application/json" {
				body = r.Body
			}
		} else {
			method = "RPC"
			path = meta["Micro-Method"]
			queries = nil
			header = meta
			body = req.Body()
		}

		defer func() {
			duration = time.Now().Sub(begin).Milliseconds()

			if w, ok := rsp.(*proto.Response); ok {
				result = w.Body
			} else {
				result = rsp
			}

			logrus.WithFields(logrus.Fields{
				"ip":       ipAddr,
				"method":   method,
				"path":     path,
				"reqId":    reqId,
				"header":   header,
				"queries":  queries,
				"reqBody":  body,
			}).Info(result)
		}()

		ctx = metadata.NewContext(ctx, map[string]string{
			"Request-Id": reqId,
		})

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
