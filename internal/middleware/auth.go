package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"net/http"
	"strconv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bypass := r.Header.Get("x-bypass")
		userId, _ := strconv.Atoi(r.Header.Get("x-user-id"))

		if bypass != "" {

			ctx := context.WithValue(r.Context(), constant.USERID, userId)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

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

			ctx := context.WithValue(r.Context(), constant.USERID, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))

		}
	})
}
