package handler

import (
	"context"
	"fmt"
	"time"

	"filestore-server/common"
	"filestore-server/config"
	cfg "filestore-server/config"
	dblayer "filestore-server/db"
	proto "filestore-server/service/account/proto"
	dbProto "filestore-server/service/dbproxy/proto"
	"filestore-server/util"

	"github.com/micro/go-micro"
)

var (
	dbCli dbProto.DBProxyService
)

// User : 用于实现UserServiceHandler接口的对象
type User struct{}

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()

	// 初始化一个dbproxy服务的客户端
	dbCli = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// Signup : 处理用户注册请求
func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, res *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		res.Code = common.StatusParamInvalid
		res.Message = "注册参数无效"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + cfg.PasswordSalt))
	// 将用户信息注册到用户表中
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		res.Code = common.StatusOK
		res.Message = "注册成功"
	} else {
		res.Code = common.StatusRegisterFailed
		res.Message = "注册失败"
	}
	return nil
}

// Signin : 处理用户登录请求
func (u *User) Signin(ctx context.Context, req *proto.ReqSignin, res *proto.RespSignin) error {
	username := req.Username
	password := req.Password

	encPasswd := util.Sha1([]byte(password + config.PasswordSalt))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		res.Code = common.StatusLoginFailed
		return nil
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		res.Code = common.StatusServerError
		return nil
	}

	// 3. 登录成功, 返回token
	res.Code = common.StatusOK
	res.Token = token
	return nil
}

// UserInfo ： 查询用户信息
func (u *User) UserInfo(ctx context.Context, req *proto.ReqUserInfo, res *proto.RespUserInfo) error {
	// uInfo, _ := json.Marshal([]interface{}{req.Username})
	// dbRes, err := dbCli.ExecuteAction(context.TODO(), &dbProto.ReqExec{
	// 	Action: []*dbProto.SingleAction{
	// 		&dbProto.SingleAction{
	// 			Name:   "/user/GetUserInfo",
	// 			Params: uInfo,
	// 		},
	// 	},
	// })
	// log.Printf("dbRpcRes: %+v %+v\n", dbRes, err)

	// 查询用户信息
	user, err := dblayer.GetUserInfo(req.Username)
	if err != nil {
		res.Code = common.StatusServerError
		res.Message = "服务错误"
		return nil
	}
	// 查不到对应的用户信息
	if user == nil {
		res.Code = common.StatusUserNotExists
		res.Message = "用户不存在"
		return nil
	}

	// 3. 组装并且响应用户数据
	res.Code = common.StatusOK
	res.Username = user.Username
	res.SignupAt = user.SignupAt
	res.LastActiveAt = user.LastActiveAt
	res.Status = int32(user.Status)
	// TODO: 需增加接口支持完善用户信息(email/phone等)
	res.Email = user.Email
	res.Phone = user.Phone
	return nil
}
