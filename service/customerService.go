package service

import (
	"github.com/cbdavid14/ms-api-go-banking/domain"
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = "2"
	}

	Customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range Customers {
		response = append(response, c.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomerById(customerId int) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.FindById(customerId)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}
