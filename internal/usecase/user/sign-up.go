package user

import (
	"context"
)

func (uc UserUsecase) SignUpUsecase(ctx context.Context) {

	uc.userRepo.SignUp(ctx)
}
