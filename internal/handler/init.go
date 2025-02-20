package handler

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/entity"
	"net/http"
)

type BillingUsecases interface {
	CreateLoanUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	MakePaymentUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
}

type UserUsecases interface {
	SignUpUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	SignInUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
}

type Handler struct {
	billingUsecases BillingUsecases
	userUsecases    UserUsecases
}

func NewBillingHandler(userUsecases UserUsecases, billingUsecases BillingUsecases) *Handler {
	return &Handler{
		userUsecases:    userUsecases,
		billingUsecases: billingUsecases,
	}
}
