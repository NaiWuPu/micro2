package Weblib

import (
	"github.com/gin-gonic/gin"
	"micro2/go-micro-grpc/Services"
)

func NewGinRouter(prodService Services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService))
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("GET", "/prods", GetProdList)
		v1Group.Handle("GET", "/prods1", GetProdList1)

	}

	return ginRouter
}
