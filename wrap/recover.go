package wrap

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
	"runtime"
)

// 防止程序异常退出
func RecoverWrapHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		meta, _ := metadata.FromContext(ctx)

		defer func() {
			if err := recover(); err != nil {
				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a endpoint('%s')\n", req.Endpoint())
				logMessage += fmt.Sprintf("MetaData: %+v\n", meta)
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)

				logrus.Errorf("recover => %s", logMessage)
			}
		}()

		return fn(ctx, req, rsp)
	}
}
