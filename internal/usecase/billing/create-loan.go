package billing

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/entity"
)

func (uc BillingUsecase) CreateLoanUsecase(ctx context.Context) (*entity.ApiResponse, error) {
	uc.billingRepo.CreateLoan(ctx)
	return nil, nil
}
