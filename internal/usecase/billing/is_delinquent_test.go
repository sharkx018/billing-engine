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

func TestIsDelinquentUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBillingRepo := mock_repo.NewMockResourceRepo(ctrl)
	billingUsecase := BillingUsecase{
		billingRepo: mockBillingRepo,
	}

	ctx := context.WithValue(context.Background(), constant.USERID, 123)

	t.Run("no loans for user", func(t *testing.T) {
		mockBillingRepo.EXPECT().GetLoanByUserId(ctx, 123).Return(nil, false)

		req := httptest.NewRequest(http.MethodGet, "/is-delinquent", nil)

		response, err := billingUsecase.IsDelinquentUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)

	})

	t.Run("user with delinquent loans", func(t *testing.T) {
		loans := []store.Loan{
			{LoanID: 1, MissedPayments: 1},
			{LoanID: 2, MissedPayments: 2},
			{LoanID: 3, MissedPayments: 0},
		}
		mockBillingRepo.EXPECT().GetLoanByUserId(ctx, 123).Return(loans, true)

		req := httptest.NewRequest(http.MethodGet, "/is-delinquent", nil)

		response, err := billingUsecase.IsDelinquentUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)

	})
}
