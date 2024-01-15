package domain

import (
	"fmt"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"github.com/cbdavid14/ms-api-go-banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
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

func (d AccountRepositoryDb) FindById(accountID int) (*Account, *errs.AppError) {
	var a Account
	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"
	err := d.client.Get(&a, query, strconv.Itoa(accountID))
	if err != nil {
		logger.Error("Error while getting account information" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tran, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	insertTransaction, err := tran.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		tran.Rollback()
		logger.Error("Error while inserting new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	if t.IsWithdrawal() {
		_, err = tran.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tran.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tran.Rollback()
		logger.Error("Error while updating account balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		logger.Error("Error while commiting transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	transactionId, err := insertTransaction.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = int(transactionId)
	t.Amount = account.Amount
	return &t, nil
}
