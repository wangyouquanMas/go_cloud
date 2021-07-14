package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"awesomeProject1/gRPC/proto/love"
)

// 定义Love服务
type Love struct {
}

// 实现Love服务接口
func (l *Love) Confession(ctx context.Context, request *love.Request) (*love.Response, error) {
	resp := &love.Response{}
	resp.Result = "I love you " + request.Name
	return resp, nil
}

func main() {
	// 监听8888端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc server
	s := grpc.NewServer()

	// 注册Love服务
	love.RegisterLoveServer(s, new(Love))

	log.Println("Listen on 127.0.0.1:8888...")
	s.Serve(listen)
}