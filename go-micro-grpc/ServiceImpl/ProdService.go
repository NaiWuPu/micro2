package ServiceImpl

import (
	"context"
	"micro2/go-micro-grpc/Services"
	"strconv"
)

type ProdService struct {
}

func NewProd(id int64, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdId: id, ProdName: pname,
	}
}
func (*ProdService) GetProdsList(ctx context.Context, in *Services.ProdsRequest, out *Services.ProdListResponse) (err error) {
	models := make([]*Services.ProdModel, 0)
	var i int64
	for i = 0; i < in.Size; i++ {
		models = append(models, NewProd(100+i, "dahaoren"+strconv.Itoa(int(i))))
	}
	out.Data = models
	return err
}
