package main

import (
	"log"

	micro "github.com/micro/go-micro"

	"filestore-server/service/account/handler"
	"filestore-server/service/account/proto"
)

func main() {
	// 创建一个service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
