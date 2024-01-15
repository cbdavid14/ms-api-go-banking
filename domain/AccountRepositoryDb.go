package domain

import (
	"fmt"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"github.com/cbdavid14/ms-api-go-banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	logger.Info(fmt.Sprintf("Creating new account %s", query))
	result, err := d.client.Exec(query, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = int(id)

	return &a, nil
}
