/**
 * Created by Goland.
 * User: yan.wang5
 * Date: 2019/9/6
 */
package admin_user

import (
	"github.com/tomdong2010/go-micro-api/app/api.gateway/service/admin.user"
	"github.com/tomdong2010/go-micro-api/http"
	"github.com/valyala/fasthttp"
)

func UserLogin(ctx *fasthttp.RequestCtx) {
	var (
		username = ctx.QueryArgs().Peek("username")
		password = ctx.QueryArgs().Peek("password")
	)

	err := admin_user.LoginByUserName(ctx, username, password)
	if err != nil {
		http.ApiRet(ctx, err)
		return
	}

	http.ApiRet(ctx, http.NewSuccess(nil))
}
