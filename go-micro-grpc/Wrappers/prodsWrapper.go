package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"micro2/go-micro-grpc/Services"
	"strconv"
)

type prodWrapper struct {
	client.Client
}

func (c *prodWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()

	configA := hystrix.CommandConfig{
		Timeout: 2000,				// 超时时间则进入降级
		ErrorPercentThreshold: 2,	// 可执行次数，否则执行熔断
		RequestVolumeThreshold: 50,	// 容忍百分比还是失败则熔断，一次熔断
		SleepWindow: 5000,			// 熔断器再次开启
	}

	hystrix.ConfigureCommand("随便取个名", configA)

	return hystrix.Do(cmdName, func() error {
		return c.Client.Call(ctx, req, rsp)
	}, func(err error) error { // 降级使用
		//defaultProds(rsp)
		defaultData(rsp)
		return nil
	})
}

func NewLogWrapper(c client.Client) client.Client {
	return &prodWrapper{c}
}

func NewProd(id int64, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdId: id, ProdName: pname,
	}
}
func defaultProds(rsp interface{}) {
	res := rsp.(*Services.ProdListResponse)
	models := make([]*Services.ProdModel, 0)
	var i int64
	for i = 0; i < 0; i++ {
		models = append(models, NewProd(100+i, "dahaoren"+strconv.Itoa(int(i))))
	}
	res.Data = models
}
func defaultData(rsp interface{}) {
	switch rsp.(type) {
	case *Services.ProdListResponse:
		defaultProds(rsp)
		// 如果有不同的返回值类型则...
	//case *Services.:
	//	NewProd(100+t.ProdId, "dahaoren"+strconv.Itoa(int(t.ProdId)))
	}
}