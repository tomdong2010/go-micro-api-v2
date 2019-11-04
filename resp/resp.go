/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package resp

import (
	"context"
	"encoding/json"
	"net/http"
	"demo/proto"
)

func Return(ctx context.Context, rw *proto.Response, resp Response) error {
	if lang, ok := ctx.Value("lang").(string); ok && lang != "" {
		resp.Translate(lang)
	}

	rw.StatusCode = http.StatusOK

	if b, err := json.Marshal(resp); err != nil {
		return err
	} else {
		rw.Body = string(b)
	}

	return nil
}
