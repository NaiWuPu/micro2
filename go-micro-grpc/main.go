package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro2/go-micro-grpc/ServiceImpl"
	"micro2/go-micro-grpc/Services"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	prodService := micro.NewService(
		micro.Name("prodService.client"),
		micro.Registry(etcdReg)) // 修饰，中间件

	prodService.Init()
	_ = Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))

	_ = prodService.Run()
}
