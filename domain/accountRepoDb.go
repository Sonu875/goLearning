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
func NewAcountRepoDb(dbClient *sqlx.DB) AccountRepoDb {
	return AccountRepoDb{dbClient}
}
