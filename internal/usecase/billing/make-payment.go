package billing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
)

func (uc BillingUsecase) MakePaymentUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	// getting the user-id from auth-token
	userID, ok := ctx.Value(constant.USERID).(int)
	if !ok {
		return nil, fmt.Errorf("user is unauthorized")
	}

	// parsing the body
	var payload entity.MakePaymentRequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
	store.GlobalStore.Mu.Lock()
	defer store.GlobalStore.Mu.Unlock()

	loan, exists := uc.billingRepo.GetLoanById(ctx, payload.LoanID)
	if !exists || loan.UserID != userID {
		return nil, fmt.Errorf("invalid loan ID or user mismatch")
	}

	// adding the validations
	if payload.EMINumber < 1 || payload.EMINumber > len(loan.EMISchedule) {
		return nil, fmt.Errorf("invalid EMI number")
	}

	emiIndex := payload.EMINumber - 1

	// checking for the edge-case if any of the previous emi is not paid
	for i := range emiIndex {
		if loan.EMISchedule[i].Status != store.Paid {
			return nil, fmt.Errorf("previous EMI is not paid")
		}
	}

	// check the edge-case if the emi is already paid
	if loan.EMISchedule[emiIndex].Status == store.Paid {
		return nil, fmt.Errorf("emi already paid")
	}

	loan.EMISchedule[emiIndex].Status = store.Paid
	loan.Outstanding -= loan.EMISchedule[emiIndex].Amount
	loan.PendingPayments--

	if loan.PendingPayments > 0 {
		loan.NextPaymentDate = loan.EMISchedule[emiIndex+1].DueDate
	}

	uc.billingRepo.UpdateLoan(ctx, loan)

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":   "Payment processed successfully",
			"loan_info": store.GlobalStore.Loans[payload.LoanID],
		},
		Success: true,
	}, nil

}
