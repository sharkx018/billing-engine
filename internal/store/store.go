package store

import "sync"

type User struct {
	UserID   int    `json:"user_id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Loan struct {
	LoanID         int     `json:"loan_id"`
	UserID         int     `json:"user_id"`
	TotalAmount    float64 `json:"total_amount"`
	WeeklyPayment  float64 `json:"weekly_payment"`
	Outstanding    float64 `json:"outstanding"`
	MissedPayments int     `json:"missed_payments"`
}

// Store struct to hold global variables
type Store struct {
	Users map[string]User
	Loans map[int]Loan
	Mu    sync.Mutex
}

// Global instance of store
var GlobalStore = Store{
	Users: make(map[string]User),
	Loans: make(map[int]Loan),
}
