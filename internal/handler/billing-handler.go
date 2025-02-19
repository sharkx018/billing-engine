package handler

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/helper"
	"net/http"
)

func (h *Handler) CreateLoanHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	h.billingUsecases.CreateLoanUsecase(ctx)
	helper.WriteCustomResp(w, http.StatusOK, "Loan Created Successfully")

}

//func (h *Handler) SignUpHandlersdfsd(w http.ResponseWriter, r *http.Request) {
//
//	ctx := context.Background()
//	h.userUsecases.SignUpUsecase(ctx)
//	helper.WriteCustomResp(w, http.StatusOK, "User Created Created Successfully")
//
//}
