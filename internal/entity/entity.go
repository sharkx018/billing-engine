package entity

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("my_secret_key")

type ApiResponse struct {
	Data    interface{}          `json:"data"`
	Success bool                 `json:"success"`
	Error   *CommonErrorResponse `json:"error,omitempty"`
}

type CreateLoadRequestPayload struct {
	Principal float64 `json:"principal"`
}

type MakePaymentRequestPayload struct {
	LoanID    int `json:"loan_id"`
	EMINumber int `json:"emi_number"`
}

type CommonErrorResponse struct {
	Message string `json:"message"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
