package domain

import (
	"fmt"
	logger "go-banking/log"
	"go-banking/util/resp_error"
	"strconv"

	"github.com/jmoiron/sqlx"
)

const (
	sqlInsertTransaction = "INSERT INTO transactions (account_id,amount,transaction_type,transaction_date) values(?,?,?,?)"
	sqlUpdateWithdrawTransaction = "UPDATE accounts SET amount = amount - ? where account_id = ?"
	sqlUpdateDeposittransaction = "UPDATE accounts SET amount = amount + ? where account_id = ?"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func NewTransactionRepositoryDB(client *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{client: client}
}

func (d TransactionRepositoryDb) MakeTransaction(t Transaction) (*Transaction, *resp_error.AppError) {
	sqlUpdate := "update from transactions set amount = ? where transaction_id = ?"

	exec, err := d.client.Exec(sqlUpdate,t.Amount,t.TransactionId)
	if err != nil {
		logger.Error(fmt.Sprintf("Error while updating transaction: %v", err.Error()))
		return nil, resp_error.NewUnexpectedError("unexpected error database")
	}

	id,err := exec.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id "+ err.Error())
		return nil, resp_error.NewUnexpectedError("unexpected error database")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t,nil
}

func (d TransactionRepositoryDb) SaveTransaction(t Transaction, accountAmount float64) (*Transaction, *resp_error.AppError) {
	//starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: "+ err.Error())
		return nil, resp_error.NewUnexpectedError("Unexpected database error")
	}
	
	//inserting bank account transaction
	result, _ := tx.Exec(sqlInsertTransaction,t.AccountId,t.Amount,t.TransactionType,t.TransactionDate)

	if t.TransactionType == "withdrawal" {
		_, err = tx.Exec(sqlUpdateWithdrawTransaction,t.Amount,t.AccountId)
		t.Amount = accountAmount - t.Amount
	} else {
		_, err = tx.Exec(sqlUpdateDeposittransaction,t.Amount,t.AccountId)
		t.Amount = t.Amount + accountAmount
	}

	// in case of error rollback, and change from both the tables will be retrieved
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: "+ err.Error())
		return nil, resp_error.NewUnexpectedError("unexpected database error")
	}

	//commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank acount: "+ err.Error())
		return nil , resp_error.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last inserted id: "+ err.Error())
		return nil, resp_error.NewUnexpectedError("unexpected database error")
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	return &t, nil
}