package repo

import (
	"context"
	"errors"
	"github.com/sharkx018/billing-engine/internal/logger"
	"github.com/sharkx018/billing-engine/internal/store"
	"golang.org/x/crypto/bcrypt"
)

func (r *ResourceRepository) RegisterUser(ctx context.Context, user store.User) (int, error) {

	logger.LogInfo(ctx, "User Sign Up Repo log")

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
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

func (r *ResourceRepository) GetUserByMobile(ctx context.Context, mobile string) (store.User, bool) {

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
	store.GlobalStore.Mu.Lock()
	storedUser, exists := store.GlobalStore.Users[mobile]
	store.GlobalStore.Mu.Unlock()

	return storedUser, exists
}
