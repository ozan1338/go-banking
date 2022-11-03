package domain

import (
	"database/sql"
	"fmt"
	logger "go-banking/log"
	"go-banking/util/resp_error"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer,*resp_error.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql)
	}
	

	if err != nil {
		logger.Error("Error While Quering customers table " + err.Error())
		return nil, resp_error.NewUnexpectedError(fmt.Sprintf("Error While Quering customers table %v", err.Error()))
	}
	
	return customers,nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *resp_error.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	// row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, resp_error.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error While Scanning Customer")
			return nil, resp_error.NewUnexpectedError("unexpected database error")
		}
	}

	return &c,nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB{
	return CustomerRepositoryDB{client: dbClient}
}