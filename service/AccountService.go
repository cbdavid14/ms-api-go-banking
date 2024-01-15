package service

import (
	"github.com/cbdavid14/ms-api-go-banking/domain"
	"github.com/cbdavid14/ms-api-go-banking/dto"
	"github.com/cbdavid14/ms-api-go-banking/errs"
	"strings"
	"time"
)

type AccountService interface {
	Save(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepositoryDb
}

func NewAccountService(repository domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{repository}
}

func (s DefaultAccountService) Save(req dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerID:  req.CustomerId,
		OpeningDate: time.Now(),
		AccountType: strings.ToLower(req.AccountType),
		Amount:      req.Amount,
		Status:      1,
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToAccountResponseDto()
	return &response, nil
}
