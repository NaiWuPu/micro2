package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
	"microProto/Models"

	myHttp "github.com/micro/go-plugins/client/http"
)

func callApi2(s selector.Selector)  {
	myClient := myHttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"))
	req := myClient.NewRequest("dahaoren", "/v1/user",
		Models.ProdsRequest{
		Size: 6,
		})

	var resp Models.ProdListResponse
	err := myClient.Call(context.Background(), req, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetData())
	//fmt.Println(resp.Data)

}

func main() {
	// 查询 etcd 认证
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.Random))
	callApi2(mySelector)
}