package service

import (
	"go-banking/domain"
	"go-banking/util/resp_error"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *resp_error.AppError)
	GetCustomer(string) (*domain.Customer, *resp_error.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *resp_error.AppError) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *resp_error.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) (DefaultCustomerService) {
	return DefaultCustomerService{repository}
}

