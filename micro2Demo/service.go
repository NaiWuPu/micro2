package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	pb "micro2/proto"
)

type Hello struct {
}

func (h Hello) HelloName(ctx context.Context, in *pb.HelloRequest, out *pb.HelloResponse) (err error) {
	fmt.Println("recived data", in.Id)
	out.Name = "欧麦有啊 心得路"
	return nil
}

func main() {

	etcdReg := etcd.NewRegistry(registry.Addrs("180.76.233.214:2379")) //注册服务到etcd中

	service := micro.NewService(
		micro.Name("hello"),
		micro.Address(":8001"),
		micro.Registry(etcdReg))

	service.Init()

	if err := pb.RegisterHelloHandler(service.Server(), new(Hello)); err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
