package domain

import appError "github.com/Sonu875/goLearning/Errors"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepo interface {
	Save(Account) (*Account, *appError.AppError)
}
