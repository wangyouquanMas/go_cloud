package main

import (
	"fmt"

	micro "github.com/micro/go-micro"

	cfg "filestore-server/service/upload/config"
	upProto "filestore-server/service/upload/proto"
	"filestore-server/service/upload/route"
	upRpc "filestore-server/service/upload/rpc"
)

var done chan bool

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.upload"), // 服务名称
	)
	service.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startApiService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	// rpc 服务
	go startRpcService()

	// api 服务
	startApiService()

	<-done
}
