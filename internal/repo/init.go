package repo

import "context"

type ResourceRepo interface {
	CreateLoan(ctx context.Context)
	SignUp(ctx context.Context)
}

type ResourceRepository struct{}

func NewResource() *ResourceRepository {
	return &ResourceRepository{}
}
