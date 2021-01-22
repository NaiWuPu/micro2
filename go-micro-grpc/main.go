package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro2/go-micro-grpc/ServiceImpl"
	"micro2/go-micro-grpc/Services"
)

// 服务端注册了 prodService.clien 负责返回接收值
func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	// 注册服务
	prodService := micro.NewService(
		micro.Name("prodService.client"),
		micro.Registry(etcdReg))

	prodService.Init()
	_ = Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))

	_ = prodService.Run()
}
