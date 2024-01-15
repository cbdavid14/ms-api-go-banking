package domain

import (
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"time"
)

const (
	TRANSACTIONTYPEWITHDRAWAL = "withdrawal"
	TRANSACTIONTYPEDEPOSIT    = "deposit"
)

type Transaction struct {
	TransactionId   int       `db:"transaction_id"`
	AccountId       int       `db:"account_id"`
	Amount          float64   `db:"amount"`
	TransactionType string    `db:"transaction_type"`
	TransactionDate time.Time `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == TRANSACTIONTYPEWITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
