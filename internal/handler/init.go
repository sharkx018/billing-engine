package handler

import "context"

type BillingUsecases interface {
	CreateLoanUsecase(ctx context.Context)
}

type BillingHandler struct {
	billingUsecases BillingUsecases
}

func NewBillingHandler(billingUsecases BillingUsecases) *BillingHandler {
	return &BillingHandler{
		billingUsecases: billingUsecases,
	}
}
