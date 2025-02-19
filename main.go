package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/handler"
	"github.com/sharkx018/billing-engine/internal/middleware"
	"github.com/sharkx018/billing-engine/internal/repo"
	"github.com/sharkx018/billing-engine/internal/usecase/billing"
	"github.com/sharkx018/billing-engine/internal/usecase/user"
	"net/http"
)

func main() {

	fmt.Println("Billing Engine!")

	dataSourceRepo := repo.NewResource()

	billingUsecase := billing.NewBillingUsecase(dataSourceRepo)
	userUsecase := user.NewUserUsecase(dataSourceRepo)

	handler := handler.NewBillingHandler(userUsecase, billingUsecase)

	router := chi.NewRouter()

	// user-route
	router.Post("/sign-up", handler.SignUpHandler)
	router.Post("/sign-in", handler.SignInHandler)

	// billing-route
	router.With(middleware.AuthMiddleware).Post("/create-loan", handler.CreateLoanHandler)

	fmt.Printf("Billing Server Started at port %s\n", constant.ConfigPort)
	err := http.ListenAndServe(constant.ConfigPort, router)
	if err != nil {
		fmt.Println("Error while starting the billing server", err.Error())
	}

}
