package handler

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/helper"
	"net/http"
)

func (h *BillingHandler) CreateLoanHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	h.billingUsecases.CreateLoanUsecase(ctx)
	helper.WriteCustomResp(w, http.StatusOK, "Loan Created Successfully")

}
