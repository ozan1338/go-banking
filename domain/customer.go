package domain

import "go-banking/util/resp_error"

type Customer struct {
	Id          string	`db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string	`db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *resp_error.AppError)
	ById(string) (*Customer, *resp_error.AppError)
}