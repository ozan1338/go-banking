package service

import (
	"go-banking/domain"
	"go-banking/util/resp_error"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *resp_error.AppError)
	GetCustomer(string) (*domain.Customer, *resp_error.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *resp_error.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *resp_error.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) (DefaultCustomerService) {
	return DefaultCustomerService{repository}
}

