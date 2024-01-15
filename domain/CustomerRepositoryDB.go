package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"github.com/cbdavid14/ms-api-go-banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:password@/banking")
	if err != nil {
		logger.Error("Error connect bd " + err.Error())
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	findAll := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	if status != "" {
		findAll = fmt.Sprintf("%s WHERE status = %s", findAll, status)
	}
	logger.Debug(findAll)
	rows, err := d.client.Query(findAll)
	if err != nil {
		logger.Error("Error while querying customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Error while querying customers table" + err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Error while scanning customers table" + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(customerId int) (*Customer, *errs.AppError) {
	findById := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	row := d.client.QueryRow(findById, customerId)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Customer not found")
		}

		logger.Error("Error while scanning customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}
