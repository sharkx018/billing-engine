package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (uc UserUsecase) SignInUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	var user store.User
	json.NewDecoder(r.Body).Decode(&user)

	// checking if the user exists or not
	storedUser, exists := uc.userRepo.GetUserByMobile(ctx, user.Mobile)

	// if user does not exist or password is wrong return error
	if !exists || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// create the jwt-token
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &entity.Claims{
		UserID: storedUser.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(constant.JwtKey)

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message": "User logged in successfully",
			"userId":  storedUser.UserID,
			"token":   tokenString,
		},
		Success: true,
	}, nil

}
