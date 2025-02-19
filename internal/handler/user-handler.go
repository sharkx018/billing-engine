package handler

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/helper"
	"net/http"
)

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.userUsecases.SignUpUsecase(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.userUsecases.SignInUsecase(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}
