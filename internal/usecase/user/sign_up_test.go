package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	mock_repo "github.com/sharkx018/billing-engine/internal/repo/mock-repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUpUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockResourceRepo(ctrl)
	userUsecase := UserUsecase{
		userRepo: mockUserRepo,
	}

	ctx := context.Background()

	t.Run("successful sign up", func(t *testing.T) {
		user := store.User{
			UserID:   123,
			Mobile:   "9876543210",
			Password: "password",
		}
		mockUserRepo.EXPECT().RegisterUser(ctx, user).Return(123, nil)

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(reqBody))

		response, err := userUsecase.SignUpUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
	})

	t.Run("registration error", func(t *testing.T) {
		user := store.User{
			Mobile:   "9876543210",
			Password: "password",
		}
		mockUserRepo.EXPECT().RegisterUser(ctx, user).Return(0, errors.New("registration failed"))

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewReader(reqBody))

		response, err := userUsecase.SignUpUsecase(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "registration failed", err.Error())
	})
}
