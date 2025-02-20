package repo

import (
	"context"
	"github.com/sharkx018/billing-engine/internal/constant"
	"github.com/sharkx018/billing-engine/internal/entity"
	"github.com/sharkx018/billing-engine/internal/store"
	"time"
)

func (b *ResourceRepository) CreateLoan(ctx context.Context, userID int, payload entity.CreateLoadRequestPayload) (store.Loan, error) {
	store.GlobalStore.Mu.Lock()
	loanID := len(store.GlobalStore.Loans) + 1
	principal := payload.Principal
	interestRate := constant.InterestRate
	totalWeeks := constant.TotalWeeks

	totalAmount := principal + (principal * interestRate)
	weeklyPayment := totalAmount / float64(totalWeeks)
	emiSchedule := make([]store.EMI, totalWeeks)

	for i := 0; i < totalWeeks; i++ {
		dueDate := time.Now().Add(time.Duration(i*7) * 24 * time.Hour).Format("2006-01-02")
		emiSchedule[i] = store.EMI{
			WeekNumber: i + 1,
			DueDate:    dueDate,
			Amount:     weeklyPayment,
			Status:     store.Pending,
		}
	}

	store.GlobalStore.Loans[loanID] = store.Loan{
		LoanID:          loanID,
		UserID:          userID,
		Principal:       principal,
		Interest:        interestRate,
		TotalAmount:     totalAmount,
		WeeklyPayment:   weeklyPayment,
		Outstanding:     totalAmount,
		MissedPayments:  0,
		NextPaymentDate: emiSchedule[0].DueDate,
		PendingPayments: totalWeeks,
		EMISchedule:     emiSchedule,
	}
	store.GlobalStore.Mu.Unlock()

	return store.GlobalStore.Loans[loanID], nil
}

func (b *ResourceRepository) GetLoanById(ctx context.Context, loanID int) (store.Loan, bool) {

	loan, exists := store.GlobalStore.Loans[loanID]
	return loan, exists
}
func (b *ResourceRepository) GetLoanByUserId(ctx context.Context, userID int) ([]store.Loan, bool) {

	var loans []store.Loan
	for _, loan := range store.GlobalStore.Loans {
		if loan.UserID == userID {
			loans = append(loans, loan)
		}
	}

	return loans, len(loans) > 0
}

func (b *ResourceRepository) UpdateLoan(ctx context.Context, loan store.Loan) (store.Loan, error) {
	store.GlobalStore.Loans[loan.LoanID] = loan
	return loan, nil
}
