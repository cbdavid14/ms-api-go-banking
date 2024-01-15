package domain

import "github.com/cbdavid14/ms-api-go-banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(int) (*Customer, *errs.AppError)
}
