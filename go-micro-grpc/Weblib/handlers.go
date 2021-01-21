package Weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro2/go-micro-grpc/Services"
)

func GetProdList(ginCtx *gin.Context) {
	//time.Sleep(3 * time.Second)
	proService := ginCtx.Keys["prodService"].(Services.ProdService)
	//context.JSON(200, prodService.NewprodList(5))
	var prodReq = new(Services.ProdsRequest)
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
		return
	} else {
		prodRes, err := proService.GetProdsList(context.Background(), prodReq)
		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
			return
		}
		//prod := new(ServiceImpl.ProdService)
		//_ = prod.GetProdsList(context.Background(), prodReq, prodRes)
		//prodRes.Data =
		ginCtx.JSON(200, gin.H{
			"data": prodRes.GetData(),
		})
	}
}
