package store

import (
	"sync"
)

type User struct {
	UserID   int    `json:"user_id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Loan struct {
	LoanID          int     `json:"loan_id"`
	UserID          int     `json:"user_id"`
	Principal       float64 `json:"principal"`
	Interest        float64 `json:"interest"`
	TotalAmount     float64 `json:"total_amount"`
	WeeklyPayment   float64 `json:"weekly_payment"`
	Outstanding     float64 `json:"outstanding"`
	MissedPayments  int     `json:"missed_payments"`
	NextPaymentDate string  `json:"next_payment_date"`
	PendingPayments int     `json:"pending_payments"`
	EMISchedule     []EMI   `json:"emi_schedule"`
}

type EMI struct {
	WeekNumber int       `json:"week_number"`
	DueDate    string    `json:"due_date"`
	Amount     float64   `json:"amount"`
	Status     EMIStatus `json:"status"`
}

const (
	Pending EMIStatus = "pending"
	Paid    EMIStatus = "paid"
	Missed  EMIStatus = "missed"
)

type EMIStatus string

// Store struct to hold global variables
type Store struct {
	Users map[int]User
	Loans map[int]Loan
	Mu    sync.Mutex
}

// Global instance of store
var GlobalStore = Store{
	Users: make(map[int]User),
	Loans: make(map[int]Loan),
}
