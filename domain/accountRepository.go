package domain

import (
	"go-banking/util/resp_error"
	"strconv"

	logger "go-banking/log"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *resp_error.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?,?,?,?,?)"
	
	exec, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating account "+ err.Error())
		return nil, resp_error.NewUnexpectedError("unexcpected error database")
	}

	id, err := exec.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id "+ err.Error())
		return nil, resp_error.NewUnexpectedError("unexcpected error database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}