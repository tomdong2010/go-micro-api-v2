/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package handler

import (
	"context"
	"demo/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.UserServerRequest, rsp *proto.UserServerResponse) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

