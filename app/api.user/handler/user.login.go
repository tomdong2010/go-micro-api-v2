/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/6
 */
package handler

import (
	"demo/app/api.user/service"
	"demo/resp"
	"github.com/kataras/iris"
)

func ActionUsers(ctx iris.Context) {

	rsp, err := service.Hello(ctx, ctx.URLParam("name"))

	if err != nil {
		resp.Err(ctx, resp.NewSystemError("service hello exception", err))
		return
	}

	resp.Suc(ctx, resp.NewSuccess(rsp))
}
