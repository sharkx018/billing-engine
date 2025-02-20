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

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
	store.GlobalStore.Mu.Lock()
	defer store.GlobalStore.Mu.Unlock()

	// get all the loans for that user
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
		// filter out the active loans and sum up the total outstanding amount
		if loan.PendingPayments != 0 && loan.Outstanding > 0 {
			activeLoans = append(activeLoans, loan)
			totalOutstanding += loan.Outstanding
		}
	}

	// return the response
	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":           "Loan info is fetched for the user",
			"total_outstanding": totalOutstanding,
			"active_loans":      activeLoans,
		},
		Success: true,
	}, nil

}
