/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/6
 */
package handler

import (
	"context"
	"demo/proto"
	"demo/resp"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"net/http"
)

type Login struct {

}

func (l *Login) Phone(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	return resp.Return(ctx, rsp, resp.NewSuccess("phone"))
}

func (l *Login) Email(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	greeter := proto.NewGreeterService("demo.srv.user", client.DefaultClient)
	res, _ := greeter.Hello(metadata.NewContext(ctx, map[string]string{
		"age": "10",
	}), &proto.UserServerRequest{
		Name: "wangy",
	})

	return resp.Return(ctx, rsp, resp.NewSuccess(res))
}

func (l *Login) Username(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.StatusCode = http.StatusOK
	rsp.Body = "username"
	return nil
}

func (l *Login) Google(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.StatusCode = http.StatusOK
	rsp.Body = "google"
	return nil
}

func (l *Login) Facebook(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.StatusCode = http.StatusOK
	rsp.Body = "facebook"
	return nil
}

func (l *Login) Twitter(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.StatusCode = http.StatusOK
	rsp.Body = "twitter"
	return nil
}

