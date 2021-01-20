package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"microProto/Helper"
	"microProto/prodService"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("8.136.142.17:2379"))

	ginRouter := gin.Default()

	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/user", func(context *gin.Context) {
			var pr Helper.ProdsRequest
			err := context.Bind(&pr)

			if err != nil || pr.Size <= 0{
				fmt.Println(err)
				pr = Helper.ProdsRequest{Size: 2}
			}
			//context.JSON(200, prodService.NewprodList(5))
			context.JSON(200, gin.H{
				"data":prodService.NewprodList(pr.Size),
			})
		})
	}

	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "news api")
	})

	server := web.NewService(
		web.Name("dahaoren"),
		web.Registry(etcdReg),
		//web.Address("127.0.0.1:8001"),	or go run main.go --server_address :8001
		web.Handler(ginRouter))

	server.Init()
	server.Run()
}
