package repo

import (
	"context"
	"errors"
	"github.com/sharkx018/billing-engine/internal/logger"
	"github.com/sharkx018/billing-engine/internal/store"
	"golang.org/x/crypto/bcrypt"
)

func (r *ResourceRepository) SignUp(ctx context.Context, user store.User) (int, error) {

	logger.LogInfo(ctx, "User Sign Up Repo log")

	store.GlobalStore.Mu.Lock()
	defer store.GlobalStore.Mu.Unlock()

	if _, exists := store.GlobalStore.Users[user.Mobile]; exists {
		return -1, errors.New("user already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.UserID = len(store.GlobalStore.Users) + 1
	user.Password = string(hashedPassword)

	store.GlobalStore.Users[user.Mobile] = user

	return user.UserID, nil
}
