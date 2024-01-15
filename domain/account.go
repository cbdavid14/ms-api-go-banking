package domain

import (
	"encoding/json"
	"github.com/cbdavid14/ms-api-go-banking/constants"
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"time"
)

type Account struct {
	AccountId   int       `db:"account_id"`
	CustomerID  int       `db:"customer_id"`
	OpeningDate time.Time `db:"opening_date"`
	AccountType string    `db:"account_type"`
	Amount      float64   `db:"amount"`
	Status      uint8     `db:"status"`
}

func (a Account) UnmarshalJSON(data []byte) error {
	var aux struct {
		AccountId   int     `json:"account_id"`
		CustomerID  int     `json:"customer_id"`
		OpeningDate string  `json:"opening_date"`
		AccountType string  `json:"account_type"`
		Amount      float64 `json:"amount"`
		Status      uint8   `json:"status"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	a.AccountId = aux.AccountId
	a.CustomerID = aux.CustomerID
	loc, err := time.Parse(constants.Commons.DateFormat, aux.OpeningDate)
	if err != nil {
		return err
	}
	a.OpeningDate = loc
	a.AccountType = aux.AccountType
	a.Amount = aux.Amount
	a.Status = aux.Status
	return nil
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

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(int) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
}
