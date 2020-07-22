/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package http

import (
	"github.com/tomdong2010/go-micro-api-v2/utility/helper"
	"github.com/valyala/fasthttp"
)

func ApiRet(ctx *fasthttp.RequestCtx, r Response) {
	b, _ := helper.JsonEncode(r)
	_, _ = ctx.Write(b)
}
