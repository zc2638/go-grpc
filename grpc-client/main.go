package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc-client/proto"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 初始化客户端
	c := proto.NewHelloClient(conn)

	// 调用方法
	reqBody := new(proto.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
		grpclog.Fatalln(err)
	}

	fmt.Println(r.Message)
	grpclog.Infoln(r.Message)
}