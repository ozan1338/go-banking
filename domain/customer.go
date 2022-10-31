package domain

import "go-banking/util/resp_error"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *resp_error.AppError)
	ById(string) (*Customer, *resp_error.AppError)
}