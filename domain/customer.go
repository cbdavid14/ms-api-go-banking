package domain

import (
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(int) (*Customer, *errs.AppError)
}
