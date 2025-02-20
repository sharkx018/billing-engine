package cron

import (
	"fmt"
	"github.com/sharkx018/billing-engine/internal/store"
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
