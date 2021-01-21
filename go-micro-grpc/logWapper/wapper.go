package logWapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
)

type logWrapper struct {
	client.Client
}

func (c *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("接口调用")
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return c.Client.Call(ctx, req, rsp)
}
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
