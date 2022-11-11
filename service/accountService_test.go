package service

import (
	"go-banking/domain"
	"go-banking/dto"
	mockDomain "go-banking/mocks/domain"
	"go-banking/util/resp_error"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	mockRepo *mockDomain.MockAccountRepository
	service AccountService
)

func TestErrorWhenTheRequestNotValidated(t *testing.T) {
	//Arange
	request := dto.NewAccountRequest{
		CustomerId: "100",
		AccountType: "saving",
		Amount: 0,
	}

	service := NewAccountService(nil)

	//Act
	_, appErr := service.NewAccount(request)

	//Assert
	if appErr == nil {
		t.Error("failed while testing a new account validation")
	}
}

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = mockDomain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestErrorWhenANewAccountCannotBeCreated(t *testing.T) {
	//Arange

	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, resp_error.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func TestNoError(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()

	
	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	
	accountWithdID := account
	accountWithdID.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithdID, nil)
	
	//Act
	newAccount, err := service.NewAccount(req)

	//Assert
	if err != nil {
		t.Error("test failed while creating new account")
	}

	if newAccount.AccountId != accountWithdID.AccountId {
		t.Error("test failed while matching new account id")
	}
}