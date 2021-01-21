package Weblib

import (
	"github.com/gin-gonic/gin"
	"micro2/go-micro-grpc/Services"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodService"] = prodService
		context.Next()
	}
}
