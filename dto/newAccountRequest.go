package dto

import (
	"go-banking/util/resp_error"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *resp_error.AppError {
	if r.Amount < 5000 {
		return resp_error.NewValidationError("To open a new account you need to deposit at least 5000")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return resp_error.NewValidationError("Account Type should be checking or saving")
	}

	return nil
}
