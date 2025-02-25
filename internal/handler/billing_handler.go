package handler

import (
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/helper"
	"net/http"
)

func (h *Handler) CreateLoanHandler(w http.ResponseWriter, r *http.Request) {

	// calling the usecase for the business logic
	ctx := r.Context()
	response, err := h.billingUsecases.CreateLoanUsecase(ctx, r)

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

func (h *Handler) MakePaymentHandler(w http.ResponseWriter, r *http.Request) {

	// calling the usecase for the business logic
	ctx := r.Context()
	response, err := h.billingUsecases.MakePaymentUsecase(ctx, r)

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

func (h *Handler) IsDelinquentHandler(w http.ResponseWriter, r *http.Request) {

	// calling the usecase for the business logic
	ctx := r.Context()
	response, err := h.billingUsecases.IsDelinquentUsecase(ctx, r)

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

func (h *Handler) GetOutstandingHandler(w http.ResponseWriter, r *http.Request) {

	// calling the usecase for the business logic
	ctx := r.Context()
	response, err := h.billingUsecases.GetOutstandingUsecase(ctx, r)

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
