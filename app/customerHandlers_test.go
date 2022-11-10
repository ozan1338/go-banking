package app

import (
	"go-banking/dto"
	"go-banking/mocks/service"
	"go-banking/util/resp_error"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	ch CustomerHandlers
	mockService *service.MockCustomerService
)

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomer)
	return func ()  {
		router = nil
		defer ctrl.Finish()
	}
}

func TestGetAllCustomersOK(t *testing.T) {
	//AAA
	//Arange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "100", Name: "ozan", City: "Jakrta", Zipcode: "123", DateOfBirth: "2000-03-22", Status: "1"},
		{Id: "101", Name: "fadhil", City: "Malang", Zipcode: "124", DateOfBirth: "1999-11-28", Status: "1"},
		{Id: "102", Name: "alfi", City: "Jakrta", Zipcode: "125", DateOfBirth: "1999-12-28", Status: "0"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	

	request, _ :=http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed testing the status code")
	}
}

func TestGetAllCustomersError(t *testing.T) {
	//AAA
	//Arange
	teardown := setup(t)
	defer teardown()
	
	mockService.EXPECT().GetAllCustomer("").Return(nil, resp_error.NewUnexpectedError("some database error"))
	

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomer)

	request, _ :=http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed testing the status code")
	}
}