package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"micro-hello/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}


func main() {
	registry := etcd.NewRegistry()

	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Registry(registry),
	)

	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	service.Run()
}
