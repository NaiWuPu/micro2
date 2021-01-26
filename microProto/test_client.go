package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	myHttp "github.com/micro/go-plugins/client/http"
	"log"
)

func main() {
	// 查询 etcd 认证
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.Random))
	getClient := myHttp.NewClient(client.Selector(mySelector), client.ContentType("application/json"))

	req := getClient.NewRequest("dahaoren", "/v1/user", map[string]string{})
	var rsp map[string]interface{}
	err := getClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
