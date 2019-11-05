/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package resp

import (
	"github.com/kataras/iris"
)

func Suc(ctx iris.Context, r Response) {
	_, _ = ctx.JSON(r)
}

func Err(ctx iris.Context, r Response) {
	if lang, ok := ctx.Values().Get("lang").(string); ok && lang != "" {
		r.Translate(lang)
	}

	_, _ = ctx.JSON(r)
}

