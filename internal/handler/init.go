package handler

import "context"

type BillingUsecases interface {
	CreateLoanUsecase(ctx context.Context)
}

type UserUsecases interface {
	SignUpUsecase(ctx context.Context)
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
