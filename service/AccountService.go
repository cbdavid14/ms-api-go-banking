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
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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
		OpeningDate: time.Now().UTC(),
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawal() {
		account, err := s.repo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	t := domain.Transaction{
		AccountId:       req.AccountId,
		TransactionDate: time.Now().UTC(),
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil
}
