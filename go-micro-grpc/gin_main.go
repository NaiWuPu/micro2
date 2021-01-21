package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"micro2/go-micro-grpc/Services"
	"micro2/go-micro-grpc/Weblib"
	"micro2/go-micro-grpc/logWapper"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	microService := micro.NewService(micro.Name("microDaHaoRenService"),
		micro.WrapClient(logWapper.NewLogWrapper))
	proService := Services.NewProdService("prodService", microService.Client())

	HttpServer := web.NewService(
		web.Name("httpDaHaoRenService"),
		web.Registry(etcdReg),
		web.Address(":8001"),
		//or go run main.go --server_address :8001
		web.Handler(Weblib.NewGinRouter(proService)),
		web.Metadata(map[string]string{"protocol": "http"}))

	HttpServer.Init()
	HttpServer.Run()
}
