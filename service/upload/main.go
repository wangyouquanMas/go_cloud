package main

import (
	"fmt"
	"time"

	micro "github.com/micro/go-micro"

	"filestore-server/config"
	cfg "filestore-server/service/upload/config"
	upProto "filestore-server/service/upload/proto"
	"filestore-server/service/upload/route"
	upRpc "filestore-server/service/upload/rpc"
)

func startRPCService() {
	service := micro.NewService(
		micro.Name("go.micro.service.upload"),   // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Registry(config.RegistryConsul()), // micro(v1.18) 需要显式指定consul (modifiled at 2020.04)
	)
	service.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
