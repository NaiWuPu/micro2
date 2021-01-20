package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func callApi(address string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+address+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buffer, _ :=ioutil.ReadAll(res.Body)
	return string(buffer), err
}

func main() {
	// 查询 etcd 认证
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	for {
		// 获取key
		etcdService, err := etcdReg.GetService("dahaoren")
		if err != nil {
			log.Fatal(err)
		}

		// 随机
		//next := selector.Random(etcdService)
		// 轮询
		next := selector.RoundRobin(etcdService)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Id, node.Address, node.Metadata)
		time.Sleep(500 * time.Millisecond)

		callRes, err := callApi(node.Address, "/v1/user", "POST")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(callRes)
	}
}
