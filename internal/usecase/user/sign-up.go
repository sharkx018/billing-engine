package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/logger"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
	"time"
)

func (uc UserUsecase) SignUpUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	var user store.User
	json.NewDecoder(r.Body).Decode(&user)

	userId, err := uc.userRepo.SignUp(ctx, user)
	if err != nil {
		return nil, err
	}

	logger.LogInfo(ctx, fmt.Sprintf("SignUp userId: %v", userId))

	// create the token
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &entity.Claims{
		UserID: int(userId),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(constant.JwtKey)

	return &entity.ApiResponse{
		Data: map[string]interface{}{

			"token": tokenString,
		},
	}, nil
}
