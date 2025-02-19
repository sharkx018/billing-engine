package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/handler"
	"github.com/sharkx018/billing-engine/internal/repo"
	"github.com/sharkx018/billing-engine/internal/usecase"
	"net/http"
)

func main() {

	fmt.Println("Billing Engine!")

	billingRepo := repo.NewBillingResource()

	billingUsecase := usecase.NewBillingUsecase(billingRepo)

	billingHandler := handler.NewBillingHandler(billingUsecase)

	router := chi.NewRouter()
	router.Post("/create-loan", billingHandler.CreateLoanHandler)

	fmt.Printf("Billing Server Started at port %s\n", constant.ConfigPort)
	err := http.ListenAndServe(constant.ConfigPort, router)
	if err != nil {
		fmt.Println("Error while starting the billing server", err.Error())
	}

}
