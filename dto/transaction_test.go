package dto

import (
	"net/http"
	"testing"
)

func TestShouldErrorTransactionTypeIsNotDepositOrWithdraw(t *testing.T) {
	// AAA (Arrange Act Assert)
	//Arrange
	request := NewTransactionRequest{
		TransactionType: "invalid tf type",
	}

	//Act
	appErr := request.Validate()

	//Assert
	if appErr.Message != "type of transaction should be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func TestShouldErrorTransactionAmountLessThanZero(t *testing.T) {
	//Arange
	request := NewTransactionRequest{
		TransactionType: "deposit",
		Amount: -1,
	}

	//Act
	appErr := request.Validate()

	//Assert
	if appErr.Message != "Amount can't be less than zero" {
		t.Error("Invalid message while testing amount transaction")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing amount transaction")
	}
}