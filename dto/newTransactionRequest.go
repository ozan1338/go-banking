package dto

import (
	"go-banking/util/resp_error"
)

var (
	withdrawal = "withdrawal"
	deposit = "deposit"
)

type NewTransactionRequest struct {
	AccountId string `json:"account_id"`
	CustomerId string `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r NewTransactionRequest) Validate() *resp_error.AppError {
	if r.TransactionType != withdrawal && r.TransactionType != deposit {
		return resp_error.NewValidationError("type of transaction should be withdrawal or deposit")
	}

	if r.Amount < 0 {
		return resp_error.NewValidationError("Amount can't be less than zero")
	}

	return nil
}

func (r NewTransactionRequest) IsWithdrawal() bool {
	return r.TransactionType == withdrawal
}