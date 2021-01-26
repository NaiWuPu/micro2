package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	pb "micro2/proto"
)

func main() {
	// Create a new service
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379")) //注册服务到etcd中

	service := micro.NewService(
		micro.Registry(etcdReg))

	service.Init()

	hello := pb.NewHelloService("hello", service.Client())

	// Call the greeter
	rsp, err := hello.HelloName(context.TODO(), &pb.HelloRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Name)
}
