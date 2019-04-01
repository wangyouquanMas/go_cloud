package handler

import (
	"context"

	"filestore-server/common"
	cfg "filestore-server/config"
	dblayer "filestore-server/db"
	"filestore-server/service/account/proto"
	"filestore-server/util"
)

type User struct{}

// Signup : 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		res.Code = common.StatusParamInvalid
		res.Mesaage = "注册参数无效"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + cfg.PasswordSalt))
	// 将用户信息注册到用户表中
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		res.Code = common.StatusOK
		res.Mesaage = "注册成功"
	} else {
		res.Code = common.StatusRegisterFailed
		res.Mesaage = "注册失败"
	}
	return nil
}
