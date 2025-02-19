package billing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"net/http"
)

func (uc BillingUsecase) CreateLoanUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	userID, ok := ctx.Value(constant.USERID).(int)
	if !ok {
		return nil, fmt.Errorf("user is unauthorized")
	}

	var payload entity.CreateLoadRequestPayload
	json.NewDecoder(r.Body).Decode(&payload)

	if payload.Principal == 0.0 {
		return nil, fmt.Errorf("principal is not valid")
	}

	createdLoad, err := uc.billingRepo.CreateLoan(ctx, userID, payload)
	if err != nil {
		return nil, err
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":   "Loan created successfully",
			"loan_info": createdLoad,
		},
		Success: true,
	}, nil

}
