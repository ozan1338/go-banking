package service

import (
	"go-banking/domain"
	"go-banking/dto"
	"go-banking/util/resp_error"
	"time"
)

type TransactionService interface {
	MakeTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse,*resp_error.AppError)
}

type DefaultTransactionService struct {
	repoTransaction domain.TransactionRepository
	repoAccount domain.AccountRepository
}

func NewTransactionService(repoTransaction domain.TransactionRepository, repoAccount domain.AccountRepository) DefaultTransactionService {
	return DefaultTransactionService{repoTransaction: repoTransaction,repoAccount: repoAccount}
}

func (s DefaultTransactionService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *resp_error.AppError) {
	//incoming request validation
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// account := domain.Account{}
	a, err := s.repoAccount.GetById(req.AccountId)
	if err != nil {
		return nil, err
	}

	// server side validation checking the available balance in the account
	if req.IsWithdrawal() {
		if !a.CanWithdraw(req.Amount) {
			return nil, resp_error.NewValidationError("insufficient balance in the account")
		}
	}

	// if all is well , build the domain object & save the transaction
	t := domain.Transaction{
		AccountId: req.AccountId,
		Amount: req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, err := s.repoTransaction.SaveTransaction(t,a.Amount)
	if err != nil {
		return nil, err
	}
	
	response := transaction.ConvertToDto()

	return response, nil
}