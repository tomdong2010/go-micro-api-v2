/**
 * Created by Goland.
 * User: yan.wang5<yan.wang5@transsion.com>
 * Date: 2019/11/3
 */
package handler

import (
	"bytes"
	"context"
	proto "github.com/tomdong2010/go-micro-api/proto/srv.user"
)

type LoginServer struct{}

func (s *LoginServer) LoginByUserName(ctx context.Context, req *proto.LoginByUserNameReq, rsp *proto.LoginByUserNameResp) error {
	if !bytes.Equal(req.Username, []byte("wyanlord")) {
		rsp.ErrNo = proto.LoginByUserNameResp_ERROR_USER
		rsp.ErrMsg = "user not found"
		return nil
	}

	if !bytes.Equal(req.Password, []byte("123456")) {
		rsp.ErrNo = proto.LoginByUserNameResp_ERROR_PWD
		rsp.ErrMsg = "invalid password"
		return nil
	}

	return nil
}
