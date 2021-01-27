package ServiceImpl

import (
	"context"
	services "micro2/micro/Services"
)

type UserService struct {

}

func (h *UserService) UserReg(ctx context.Context, in *services.UserModel, out *services.RegResponse) error {
	out.Message = "123"
	out.Status = "aa"
	return nil
}
