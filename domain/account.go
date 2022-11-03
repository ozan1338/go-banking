package domain

import (
	"go-banking/dto"
	"go-banking/util/resp_error"
)

type Account struct {
	AccountId  string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *resp_error.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
