package billing

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sharkx018/billing-engine/internal/constant"
	mock_repo "github.com/sharkx018/billing-engine/internal/repo/mock-repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOutstandingUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBillingRepo := mock_repo.NewMockResourceRepo(ctrl)
	billingUsecase := BillingUsecase{
		billingRepo: mockBillingRepo,
	}

	ctx := context.WithValue(context.Background(), constant.USERID, 123)

	t.Run("no loans for user", func(t *testing.T) {
		mockBillingRepo.EXPECT().GetLoanByUserId(ctx, 123).Return(nil, false)

		req := httptest.NewRequest(http.MethodGet, "/get-outstanding", nil)

		response, err := billingUsecase.GetOutstandingUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)

	})

	t.Run("with active loans", func(t *testing.T) {
		loans := []store.Loan{
			{LoanID: 1, PendingPayments: 3, Outstanding: 500.0},
			{LoanID: 2, PendingPayments: 0, Outstanding: 0.0},
			{LoanID: 3, PendingPayments: 2, Outstanding: 700.0},
		}
		mockBillingRepo.EXPECT().GetLoanByUserId(ctx, 123).Return(loans, true)

		req := httptest.NewRequest(http.MethodGet, "/get-outstanding", nil)

		response, err := billingUsecase.GetOutstandingUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)

	})
}
