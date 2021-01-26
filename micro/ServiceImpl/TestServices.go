package ServiceImpl

import (
	"context"
	services "micro2/micro/Services"
	"strconv"
)

type TestService struct {
}

func (c *TestService) Call(ctx context.Context, req *services.TestRequest, resp *services.TestResponse) error {
	resp.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}
