package usecase

import "context"

type BillingRepo interface {
	CreateLoan(ctx context.Context)
}

type BillingUsecase struct {
	billingRepo BillingRepo
}

func NewBillingUsecase(billingresource BillingRepo) *BillingUsecase {
	return &BillingUsecase{
		billingRepo: billingresource,
	}
}
