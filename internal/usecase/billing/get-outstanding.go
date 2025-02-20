package billing

import (
	"context"
	"fmt"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
)

func (uc *BillingUsecase) GetOutstandingUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	userID, ok := ctx.Value(constant.USERID).(int)
	if !ok {
		return nil, fmt.Errorf("user is unauthorized")
	}

	store.GlobalStore.Mu.Lock()
	defer store.GlobalStore.Mu.Unlock()

	loans, exists := uc.billingRepo.GetLoanByUserId(ctx, userID)

	if !exists {
		return &entity.ApiResponse{
			Data: map[string]interface{}{
				"message":           "There is no loan with this user",
				"total_outstanding": 0,
			},
			Success: true,
		}, nil
	}

	var activeLoans []store.Loan
	var totalOutstanding float64
	for _, loan := range loans {
		if loan.PendingPayments != 0 && loan.Outstanding > 0 {
			activeLoans = append(activeLoans, loan)
			totalOutstanding += loan.Outstanding
		}
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":           "Loan info is fetched for the user",
			"active_loans":      activeLoans,
			"total_outstanding": totalOutstanding,
		},
		Success: true,
	}, nil

}
