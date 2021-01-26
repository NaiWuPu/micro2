package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro2/micro/ServiceImpl"
	services "micro2/micro/Services"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	myService := micro.NewService(
		micro.Name("test"),
		micro.Address(":8001"),
		micro.Registry(etcdReg))

	myService.Init()
	_ = services.RegisterTestServiceHandler(myService.Server(), new(ServiceImpl.TestService))

	_ = myService.Run()
}
