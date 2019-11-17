package admin_user

import (
	"demo/app/api.gateway/service"
	"demo/conf"
	"demo/http"
	"demo/proto/srv.user"
	"github.com/valyala/fasthttp"
)

func LoginByUserName(ctx *fasthttp.RequestCtx, username, password []byte) http.Response {
	loginService := proto.NewLoginService(conf.APP_SRV_USER, service.NewClient())

	rsp, err := loginService.LoginByUserName(service.NewContext(ctx), &proto.LoginByUserNameReq{
		Username: username,
		Password: password,
	})

	if err != nil {
		return http.NewSystemError("login service exception.", err)
	}

	switch rsp.ErrNo {
	case proto.LoginByUserNameResp_ERROR_NIL:
		return nil
	case proto.LoginByUserNameResp_ERROR_SYS:
		return http.NewSystemError(rsp.ErrMsg)
	case proto.LoginByUserNameResp_ERROR_USER:
		return http.NewUserNotFoundError(rsp.ErrMsg)
	case proto.LoginByUserNameResp_ERROR_PWD:
		return http.NewPasswordInvalidError(rsp.ErrMsg)
	default:
		return http.NewSystemError(rsp.ErrMsg)
	}
}