package service

import (
	"go-banking/domain"
	"go-banking/dto"
	"go-banking/util/resp_error"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *resp_error.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *resp_error.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *resp_error.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse

	for _,item := range c {
		response = append(response, item.ToDto())
	}

	return response,nil
}
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *resp_error.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	
	return &response,nil
}

func NewCustomerService(repository domain.CustomerRepository) (DefaultCustomerService) {
	return DefaultCustomerService{repository}
}

