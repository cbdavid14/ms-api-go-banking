package domain

import (
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"time"
)

type Account struct {
	AccountId   int
	CustomerID  int
	OpeningDate time.Time
	AccountType string
	Amount      float64
	Status      uint8
}

func (a Account) ToAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerID,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
