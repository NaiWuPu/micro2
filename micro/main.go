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
		micro.Name("api.wuzi.com.test"),
		micro.Address("127.0.0.1:8002"),
		micro.Registry(etcdReg))

	//err := services.RegisterTestServiceHandler(myService.Server(), new(ServiceImpl.TestService))
	//fmt.Println(err)
	//err = myService.Run()
	//fmt.Println(err)

	services.RegisterUserServiceHandler(myService.Server(), new(ServiceImpl.UserService))

	myService.Run()
}
