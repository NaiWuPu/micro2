package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"micro2/go-micro-grpc/Services"
	"micro2/go-micro-grpc/Weblib"
	"micro2/go-micro-grpc/Wrappers"
)

// 客户端负责处理参数，并向下层服务端 讨要prod的数据
func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	// 增加micro中间件
	//microService := micro.NewService(micro.WrapClient(logWapper.NewLogWrapper))

	microService := micro.NewService(micro.WrapClient(Wrappers.NewLogWrapper))

	// 发现服务
	proService := Services.NewProdService("prodService.client", microService.Client())

	// 注册服务
	HttpServer := web.NewService(
		web.Name("http_卵用server"),
		web.Registry(etcdReg),
		web.Address(":8001"),
		//or go run main.go --server_address :8001
		web.Handler(Weblib.NewGinRouter(proService)),
		web.Metadata(map[string]string{"protocol": "http"}))

	HttpServer.Init()
	HttpServer.Run()
}
