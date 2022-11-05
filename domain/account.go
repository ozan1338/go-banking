package domain

import (
	"go-banking/dto"
	"go-banking/util/resp_error"
)

type Account struct {
	AccountId  string `db:"account_id"`
	CustomerId string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *resp_error.AppError)
	GetById(string) (*Account, *resp_error.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
