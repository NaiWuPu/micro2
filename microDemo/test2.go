package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"

	myHttp "github.com/micro/go-plugins/client/http"

)

func callApi2(s selector.Selector)  {
	myClient := myHttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"))
	req := myClient.NewRequest("dahaoren", "/v1/user", map[string]int{"size":4})
	var resp map[string]interface{}
	err := myClient.Call(context.Background(), req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp["data"])
}

func main() {
	// 查询 etcd 认证
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.Random))
	callApi2(mySelector)
}