package billing

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	mock_repo "github.com/sharkx018/billing-engine/internal/repo/mock-repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoanUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBillingRepo := mock_repo.NewMockResourceRepo(ctrl)
	billingUsecase := BillingUsecase{
		billingRepo: mockBillingRepo,
	}

	ctx := context.WithValue(context.Background(), constant.USERID, 123)

	t.Run("success case", func(t *testing.T) {
		payload := entity.CreateLoadRequestPayload{
			Principal: 1000.0,
		}
		requestBody, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/create-loan", bytes.NewReader(requestBody))

		mockBillingRepo.EXPECT().CreateLoan(ctx, 123, payload).Return(store.Loan{LoanID: 1, Principal: 1000.0}, nil)

		response, err := billingUsecase.CreateLoanUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
		//assert.Equal(t, "Loan created successfully", response.Data["message"])
	})

	t.Run("unauthorized user", func(t *testing.T) {
		invalidCtx := context.Background()
		req := httptest.NewRequest(http.MethodPost, "/create-loan", nil)

		response, err := billingUsecase.CreateLoanUsecase(invalidCtx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "user is unauthorized", err.Error())
	})

	t.Run("invalid principal", func(t *testing.T) {
		payload := entity.CreateLoadRequestPayload{
			Principal: 0.0,
		}
		requestBody, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/create-loan", bytes.NewReader(requestBody))

		response, err := billingUsecase.CreateLoanUsecase(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "principal is not valid", err.Error())
	})

	t.Run("repository error", func(t *testing.T) {
		payload := entity.CreateLoadRequestPayload{
			Principal: 1000.0,
		}
		requestBody, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/create-loan", bytes.NewReader(requestBody))

		mockBillingRepo.EXPECT().CreateLoan(ctx, 123, payload).Return(store.Loan{}, errors.New("repo error"))

		response, err := billingUsecase.CreateLoanUsecase(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "repo error", err.Error())
	})
}
