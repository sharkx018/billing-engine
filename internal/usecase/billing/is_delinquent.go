package billing

import (
	"context"
	"fmt"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
)

func (uc *BillingUsecase) IsDelinquentUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	// getting the user-id from the auth-token
	userID, ok := ctx.Value(constant.USERID).(int)
	if !ok {
		return nil, fmt.Errorf("user is unauthorized")
	}

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
	store.GlobalStore.Mu.Lock()
	defer store.GlobalStore.Mu.Unlock()

	loans, exists := uc.billingRepo.GetLoanByUserId(ctx, userID)

	if !exists {
		return &entity.ApiResponse{
			Data: map[string]interface{}{
				"message":       "There is no loan with this user",
				"is_delinquent": false,
			},
			Success: true,
		}, nil
	}

	var delinquent bool
	for _, loan := range loans {
		if loan.MissedPayments >= 2 {
			delinquent = true
			break
		}
	}

	// returning the response
	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":       "Loan info is fetched for the user",
			"is_delinquent": delinquent,
		},
		Success: true,
	}, nil

}
