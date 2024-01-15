package dto

import "time"

type TransactionResponse struct {
	TransactionId   int       `json:"transaction_id"`
	AccountId       int       `json:"account_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
}
