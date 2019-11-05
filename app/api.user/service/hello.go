package service

import (
	"demo/cmn"
	"demo/proto"
	"github.com/kataras/iris"
)

func Hello(ctx iris.Context, name string) (*proto.UserServerResponse, error) {
	greeter := proto.NewGreeterService(cmn.APP_NAME_PREFIX + cmn.APP_SRV_USER, microService.Client())

	return greeter.Hello(NewContext(ctx), &proto.UserServerRequest{
		Name: name,
	})
}
