package user

import (
	"github.com/sharkx018/billing-engine/internal/repo"
)

type UserUsecase struct {
	userRepo repo.ResourceRepo
}

func NewUserUsecase(dataResourceRepo repo.ResourceRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: dataResourceRepo,
	}
}
