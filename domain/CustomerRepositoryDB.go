package domain

import (
	"fmt"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"github.com/cbdavid14/ms-api-go-banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	if status != "2" {
		query = fmt.Sprintf("%s WHERE status = %s", query, status)
	}
	logger.Debug(query)
	err := d.client.Select(&customers, query)
	if err != nil {
		logger.Error("Error while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customers table" + err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(customerId int) (*Customer, *errs.AppError) {
	var c Customer
	query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	err := d.client.Get(&c, query, customerId)
	if err != nil {
		logger.Error("Error while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customers table" + err.Error())
	}
	return &c, nil
}
