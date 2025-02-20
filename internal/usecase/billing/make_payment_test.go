package billing

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	mock_repo "github.com/sharkx018/billing-engine/internal/repo/mock-repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakePaymentUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBillingRepo := mock_repo.NewMockResourceRepo(ctrl)
	billingUsecase := BillingUsecase{
		billingRepo: mockBillingRepo,
	}

	ctx := context.WithValue(context.Background(), constant.USERID, 123)

	t.Run("successful payment", func(t *testing.T) {
		payload := entity.MakePaymentRequestPayload{
			LoanID:    1,
			EMINumber: 1,
		}
		loan := store.Loan{
			LoanID:          1,
			UserID:          123,
			Outstanding:     1000.0,
			PendingPayments: 5,
			EMISchedule: []store.EMI{
				{Amount: 200.0, Status: store.Pending},
				{Amount: 200.0, Status: store.Pending},
				{Amount: 200.0, Status: store.Pending},
				{Amount: 200.0, Status: store.Pending},
				{Amount: 200.0, Status: store.Pending},
			},
		}

		mockBillingRepo.EXPECT().GetLoanById(gomock.Any(), gomock.Any()).Return(loan, true)
		mockBillingRepo.EXPECT().UpdateLoan(gomock.Any(), gomock.Any()).Return(store.Loan{}, nil)

		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/make-payment", bytes.NewReader(reqBody))

		response, err := billingUsecase.MakePaymentUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
	})
}
