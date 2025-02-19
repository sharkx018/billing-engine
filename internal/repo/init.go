package repo

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/store"
)

type ResourceRepo interface {
	CreateLoan(ctx context.Context)
	SignUp(ctx context.Context, user store.User) (int, error)
}

type ResourceRepository struct{}

func NewResource() *ResourceRepository {
	return &ResourceRepository{}
}
