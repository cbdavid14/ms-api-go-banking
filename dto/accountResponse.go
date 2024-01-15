package dto

import "time"

type AccountResponse struct {
	AccountId   int       `json:"account_id"`
	CustomerId  int       `json:"customer_id"`
	OpeningDate time.Time `json:"opening_date"`
	AccountType string    `json:"account_type"`
	Amount      float64   `json:"amount"`
	Status      uint8     `json:"status"`
}
