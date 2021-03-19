package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"micro-hello/proto"
)

func main() {
	registry := etcd.NewRegistry()

	service := micro.NewService(
		micro.Name("helloworld.client"),
		micro.Registry(registry),
	)
	service.Init()

	greeter := proto.NewGreeterService("helloworld", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{
		Name: "chris",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Msg)
}
