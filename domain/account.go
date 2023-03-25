package domain

import (
	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/dto"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type Transaction struct {
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionId   string
}

type AccountRepo interface {
	Save(Account) (*Account, *appError.AppError)
	FindAccountById(string) (*Account, *appError.AppError)
	MakeTransaction(Transaction, float64) (*Transaction, *appError.AppError)
}

func (tx Transaction) ResponseDto(balance float64) dto.TransactionResponse {
	response := dto.TransactionResponse{
		TransactionId: tx.TransactionId,
		Balance:       balance,
	}
	return response
}
