package billing

import (
	"github.com/sharkx018/billing-engine/internal/repo"
)

type BillingUsecase struct {
	billingRepo repo.ResourceRepo
}

func NewBillingUsecase(billingresource repo.ResourceRepo) *BillingUsecase {
	return &BillingUsecase{
		billingRepo: billingresource,
	}
}
