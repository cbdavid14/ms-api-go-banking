package dto

import (
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"strings"
)

type AccountRequest struct {
	CustomerId  int     `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r AccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
