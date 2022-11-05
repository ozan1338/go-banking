package domain

import (
	"go-banking/dto"
	"go-banking/util/resp_error"
)

type Transaction struct {
	TransactionId   string `db:"transaction_id"`
	AccountId       string `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
}

type TransactionRepository interface {
	MakeTransaction(Transaction) (*Transaction, *resp_error.AppError)
	SaveTransaction(Transaction,float64) (*Transaction, *resp_error.AppError)
}

func (t Transaction) ConvertToDto() *dto.NewTransactionResponse {
	return &dto.NewTransactionResponse{
		TransactionId: t.TransactionType,
		AccountId: t.AccountId,
		NewBalance: t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}