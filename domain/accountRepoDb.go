package domain

import (
	"strconv"

	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepoDb struct {
	client *sqlx.DB
}

func (d AccountRepoDb) Save(acc Account) (*Account, *appError.AppError) {
	sqlInsert := "insert into accounts (customer_id,account_type,amount,status) values($1,$2,$3,$4) RETURNING account_id"
	var lastId int
	err := d.client.Get(&lastId, sqlInsert, acc.CustomerId, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		logger.Error("Something went wrong with DB" + err.Error())
		return nil, appError.NewInternalServerError("Something went wrong with DB")
	}
	acc.AccountId = strconv.FormatInt(int64(lastId), 10)
	return &acc, nil

}

func (d AccountRepoDb) FindAccountById(accountId string) (*Account, *appError.AppError) {
	sqlSearch := "select *  from accounts where account_id=$1"
	var accountDetails []Account
	err := d.client.Select(&accountDetails, sqlSearch, accountId)
	if err != nil {
		logger.Error("Not able to find the account you looking for" + err.Error())
		return nil, appError.NewInternalServerError("Not able to find the account you looking for ")
	}

	return &accountDetails[0], nil

}

func (d AccountRepoDb) MakeTransaction(tx Transaction, balance float64) (*Transaction, *appError.AppError) {

	sqlAlter := "UPDATE accounts SET amount=$1 WHERE account_id=$2"
	_, err_alter := d.client.Exec(sqlAlter, balance, tx.AccountId)
	if err_alter != nil {
		logger.Error("Unable to update account amount" + err_alter.Error())
		return nil, appError.NewInternalServerError("Something went wrong with DB")
	}
	sqlInsert := "insert into transactions (account_id,amount,transaction_type) values($1,$2,$3) RETURNING transaction_id"
	var lastId int
	err := d.client.Get(&lastId, sqlInsert, tx.AccountId, tx.Amount, tx.TransactionType)
	if err != nil {
		logger.Error("Something went wrong with DB" + err.Error())
		return nil, appError.NewInternalServerError("Something went wrong with DB")
	}
	tx.TransactionId = strconv.FormatInt(int64(lastId), 10)
	return &tx, nil
}

func NewAcountRepoDb(dbClient *sqlx.DB) AccountRepoDb {
	return AccountRepoDb{dbClient}
}
