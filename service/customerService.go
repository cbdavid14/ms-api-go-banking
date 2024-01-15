package service

import (
	"github.com/cbdavid14/ms-api-go-banking/domain"
	"github.com/cbdavid14/ms-api-go-banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(int) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {

	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = "2"
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerById(customerId int) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(customerId)
}
