package billing

import (
	"context"
)

func (uc BillingUsecase) CreateLoanUsecase(ctx context.Context) {
	uc.billingRepo.CreateLoan(ctx)
}
