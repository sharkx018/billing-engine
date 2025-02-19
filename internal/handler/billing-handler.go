package handler

import (
	"fmt"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/helper"
	"github.com/sharkx018/billing-engine/internal/logger"
	"net/http"
)

func (h *Handler) CreateLoanHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(constant.USERID).(int)
	if !ok {
		helper.WriteCustomResp(w, http.StatusUnauthorized, "user is unauthorized")
		http.Error(w, constant.Unauthorized, http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	logger.LogInfo(ctx, "Authorised user_id", userID)

	h.billingUsecases.CreateLoanUsecase(ctx)
	helper.WriteCustomResp(w, http.StatusOK, fmt.Sprintf("Loan Created Successfully for user_id: %d", userID))

}
