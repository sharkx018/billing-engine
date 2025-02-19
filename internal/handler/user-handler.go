package handler

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/helper"
	"net/http"
)

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	h.userUsecases.SignUpUsecase(ctx)
	helper.WriteCustomResp(w, http.StatusOK, "User Created Created Successfully")

}
