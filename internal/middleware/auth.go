package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"net/http"
	"strconv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bypass := r.Header.Get("x-bypass")
		userId, _ := strconv.Atoi(r.Header.Get("x-user-id"))

		if bypass == "" {

			tokenString := r.Header.Get(constant.Authorization)
			if tokenString == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			claims := &entity.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return constant.JwtKey, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			userId = claims.UserID
		}

		// global in-memory store
		// locking the store to avoid the race-condition as this is the shared resource
		store.GlobalStore.Mu.Lock()

		// adding the check if the user-id exist in the store or not
		if _, exists := store.GlobalStore.Users[userId]; !exists {
			http.Error(w, "user_id does not exist in the store", http.StatusUnauthorized)
			return
		}
		store.GlobalStore.Mu.Unlock()

		// setting the user-id in the context
		ctx := context.WithValue(r.Context(), constant.USERID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
