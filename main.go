package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/handler"
	"github.com/sharkx018/billing-engine/internal/logger"
	"github.com/sharkx018/billing-engine/internal/middleware"
	"github.com/sharkx018/billing-engine/internal/repo"
	"github.com/sharkx018/billing-engine/internal/store"
	"github.com/sharkx018/billing-engine/internal/usecase/billing"
	"github.com/sharkx018/billing-engine/internal/usecase/user"
	"net/http"
	"time"
)

func RunCronJob() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		<-ticker.C
		store.GlobalStore.Mu.Lock()
		fmt.Println("Running daily EMI status update job...")
		currentDate := time.Now().Format("2006-01-02")
		for loanID, loan := range store.GlobalStore.Loans {
			for i, emi := range loan.EMISchedule {
				if emi.Status == store.Pending && emi.DueDate < currentDate {
					currLoan := store.GlobalStore.Loans[loanID]
					currLoan.EMISchedule[i].Status = store.Missed
					currLoan.MissedPayments++
					store.GlobalStore.Loans[loanID] = currLoan
				}
			}
		}
		store.GlobalStore.Mu.Unlock()
	}
}

func main() {

	logger.LogInfo(context.Background(), "Billing Engine! started")

	//Cron job started
	logger.LogInfo(context.Background(), "Cron job started for marking the missed payments", "")
	go RunCronJob()

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
	router.With(middleware.AuthMiddleware).Post("/make-payment", handler.MakePaymentHandler)
	router.With(middleware.AuthMiddleware).Get("/is-delinquent", handler.IsDelinquentHandler)
	router.With(middleware.AuthMiddleware).Get("/get-outstanding", handler.GetOutstandingHandler)

	fmt.Printf("Billing Server Started at port %s\n", constant.ConfigPort)
	err := http.ListenAndServe(constant.ConfigPort, router)
	if err != nil {
		fmt.Println("Error while starting the billing server", err.Error())
	}

}
