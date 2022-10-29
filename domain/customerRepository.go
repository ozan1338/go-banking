package domain

import (
	"database/sql"
	"go-banking/util/resp_error"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_driver = "DB_DRIVER"
	db_source = "DB_SOURCE"
)

var (
	DBDriver = os.Getenv(db_driver)
	DBSource = os.Getenv(db_source)
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer,error) {

	findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
	
	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error While Quering customers table" + err.Error())
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error While Scanning customers" + err.Error())
		}
		customers = append(customers, c)
	}

	return customers,nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *resp_error.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, resp_error.NewNotFoundError("customer not found")
		} else {
			log.Println("Error While Scanning Customer")
			return nil, resp_error.NewUnexpectedError("unexpected database error")
		}
	}

	return &c,nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB{
	client, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client: client}
}