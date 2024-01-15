package domain

import "github.com/cbdavid14/ms-api-go-banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(int) (*Customer, *errs.AppError)
}
