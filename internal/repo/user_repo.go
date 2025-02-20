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

	for _, existingUser := range store.GlobalStore.Users {
		if existingUser.Mobile == user.Mobile {
			return -1, errors.New("user already exists")
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.UserID = len(store.GlobalStore.Users) + 1
	user.Password = string(hashedPassword)

	store.GlobalStore.Users[user.UserID] = user

	return user.UserID, nil
}

func (r *ResourceRepository) GetUserByMobile(ctx context.Context, mobile string) (store.User, bool) {

	// global in-memory store
	// locking the store to avoid the race-condition as this is the shared resource
	store.GlobalStore.Mu.Lock()
	var user store.User
	var exists = false

	for _, existingUser := range store.GlobalStore.Users {
		if existingUser.Mobile == mobile {
			user = existingUser
			exists = true
		}
	}

	store.GlobalStore.Mu.Unlock()
	return user, exists
}
