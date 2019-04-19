package db

import (
	"context"
	"encoding/json"

	dbProto "filestore-server/service/dbproxy/proto"

	"github.com/micro/go-micro"
)

var (
	dbCli dbProto.DBProxyService
)

func init() {
	service := micro.NewService()
	// 初始化， 解析命令行参数等
	service.Init()
	// 初始化一个dbproxy服务的客户端
	dbCli = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

func GetUserInfo(username string) (*dbProto.RespExec, error) {
	uInfo, _ := json.Marshal([]interface{}{username})
	return dbCli.ExecuteAction(context.TODO(), &dbProto.ReqExec{
		Action: []*dbProto.SingleAction{
			&dbProto.SingleAction{
				Name:   "/user/GetUserInfo",
				Params: uInfo,
			},
		},
	})
}

func UserSignup(username, encPasswd string) (*dbProto.RespExec, error) {
	// TODO: 完善逻辑
	// uInfo, _ := json.Marshal([]interface{}{username})
	return nil, nil
}
