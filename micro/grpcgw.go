package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	ServiceGW "micro2/micro/ServiceGW"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gRpcEndPoint := "localhost:8001"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} //不使用证书校验
	err := ServiceGW.RegisterTestServiceHandlerFromEndpoint(ctx, mux, gRpcEndPoint, opts) //gRpcEndPoint在这里的作用是当有请求来到9000端口会转发给8001端口
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":9000", mux)//通过postman访问9000端口服务会转发给8001端口的rpc服务
}