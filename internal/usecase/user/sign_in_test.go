package user

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	mock_repo "github.com/sharkx018/billing-engine/internal/repo/mock-repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignInUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repo.NewMockResourceRepo(ctrl)
	userUsecase := UserUsecase{
		userRepo: mockUserRepo,
	}

	ctx := context.Background()

	t.Run("successful sign in", func(t *testing.T) {
		password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		storedUser := store.User{
			UserID:   123,
			Mobile:   "9876543210",
			Password: string(password),
		}
		mockUserRepo.EXPECT().GetUserByMobile(ctx, storedUser.Mobile).Return(storedUser, true)

		user := store.User{
			Mobile:   "9876543210",
			Password: "password",
		}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader(reqBody))

		response, err := userUsecase.SignInUsecase(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
	})

	t.Run("invalid credentials", func(t *testing.T) {
		mockUserRepo.EXPECT().GetUserByMobile(ctx, "9876543210").Return(store.User{}, false)

		user := store.User{
			Mobile:   "9876543210",
			Password: "wrongpassword",
		}
		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/sign-in", bytes.NewReader(reqBody))

		response, err := userUsecase.SignInUsecase(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "invalid credentials", err.Error())
	})
}
