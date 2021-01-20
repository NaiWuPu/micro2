package prodService

import "strconv"

type ProdModel struct {
	ProdId   int    `json:"pid"`
	ProdName string `json:"pname"`
}

func NewProd(id int, pname string) *ProdModel {
	return &ProdModel{
		id, pname,
	}
}

func NewprodList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)

	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prodName"+strconv.Itoa(100+i)))
	}

	return ret
}
