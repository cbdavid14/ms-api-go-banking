package dto

import (
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"time"
)

const (
	TRANSACTIONTYPEWITHDRAWAL = "withdrawal"
	TRANSACTIONTYPEDEPOSIT    = "deposit"
)

type TransactionRequest struct {
	AccountId       int       `json:"account_id"`
	CustomerId      int       `json:"customer_id"`
	TransactionDate time.Time `json:"transaction_date"`
	TransactionType string    `json:"transactions_type"`
	Amount          float64   `json:"amount"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	if r.TransactionType != TRANSACTIONTYPEDEPOSIT && r.TransactionType != TRANSACTIONTYPEWITHDRAWAL {
		return errs.NewValidationError("Transaction type should be deposit or withdrawal")
	}
	return nil
}

func (r TransactionRequest) IsWithdrawal() bool {
	return r.TransactionType == TRANSACTIONTYPEWITHDRAWAL
}

func (r TransactionRequest) IsDeposit() bool {
	return r.TransactionType == TRANSACTIONTYPEDEPOSIT
}
