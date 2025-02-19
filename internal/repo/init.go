package repo

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
)

type ResourceRepo interface {
	CreateLoan(ctx context.Context, id int, payload entity.CreateLoadRequestPayload) (store.Loan, error)
	RegisterUser(ctx context.Context, user store.User) (int, error)
	GetUserByMobile(ctx context.Context, mobile string) (store.User, bool)
}

type ResourceRepository struct{}

func NewResource() *ResourceRepository {
	return &ResourceRepository{}
}
