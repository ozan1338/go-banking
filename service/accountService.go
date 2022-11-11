package service

import (
	"go-banking/domain"
	"go-banking/dto"
	"go-banking/util/resp_error"
)
const dbTSLayout = "2006-01-02 15:04:05"
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *resp_error.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *resp_error.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	
	newAccount , err := s.repo.Save(account)
	if err != nil {
		return nil,err
	}

	response  := newAccount.ToNewAccountResponseDto()

	return &response,nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}