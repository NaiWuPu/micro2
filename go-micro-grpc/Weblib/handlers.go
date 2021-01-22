package Weblib

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"micro2/go-micro-grpc/Services"
	"strconv"
	"time"
)

func GetProdList1(ginCtx *gin.Context) {
	time.Sleep(time.Second)
	proService := ginCtx.Keys["prodService"].(Services.ProdService)
	//context.JSON(200, prodService.NewprodList(5))
	var prodReq = new(Services.ProdsRequest)
	var prodRes = new(Services.ProdListResponse)
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		return
	}
	prodRes, _ = proService.GetProdsList(context.Background(), prodReq)
	ginCtx.JSON(200, gin.H{
		"data": prodRes.GetData(),
	})

}

func GetProdList(ginCtx *gin.Context) {
	proService := ginCtx.Keys["prodService"].(Services.ProdService)
	//context.JSON(200, prodService.NewprodList(5))
	var prodReq = new(Services.ProdsRequest)
	var prodRes = new(Services.ProdListResponse)
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		return
	} else {
		// 熔断代码改造
		{
			// 第一部，配置config
			configA := hystrix.CommandConfig{
				Timeout: 1000,
			}
			// 第二步，配置command
			hystrix.ConfigureCommand("suibian", configA)
			// 执行DO方法
			err := hystrix.Do("suibian", func() error {
				prodRes, err = proService.GetProdsList(context.Background(), prodReq)
				return err
			}, func(err error) error { // 降级使用
				prodRes, err = defaultProds()
				return err
			})

			if err != nil {
				ginCtx.JSON(500, gin.H{"status": err.Error()})
			} else {
				ginCtx.JSON(200, gin.H{
					"data": prodRes.GetData(),
				})
			}
		}
	}
}

func NewProd(id int64, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdId: id, ProdName: pname,
	}
}
func defaultProds() (*Services.ProdListResponse, error) {
	models := make([]*Services.ProdModel, 0)
	var i int64
	for i = 0; i < 0; i++ {
		models = append(models, NewProd(100+i, "dahaoren"+strconv.Itoa(int(i))))
	}
	res := &Services.ProdListResponse{}
	res.Data = models
	return res, nil
}
