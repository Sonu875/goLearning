package dto

import (
	"strings"

	appError "github.com/Sonu875/goLearning/Errors"
)

type AccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}
type AccountResponse struct {
	AccountId string `json:"account_id"`
}

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}

func (rq AccountRequest) Validate() *appError.AppError {
	if rq.Amount < 5000 {
		return appError.NewValidation("To open an account amount should be greater than 5000")
	}
	if strings.ToLower(rq.AccountType) != "saving" && strings.ToLower(rq.AccountType) != "current" {
		return appError.NewValidation("Account can be of two type current or savings")
	}
	return nil
}

func (rq TransactionRequest) Validate() *appError.AppError {
	if strings.ToLower(rq.TransactionType) != "withdrawal" && strings.ToLower(rq.TransactionType) != "deposit" {
		return appError.NewValidation("Transaction type can only be desposit or withdrawal ")
	}
	if rq.Amount < 0 {
		return appError.NewValidation("Amount should be greater than 0")
	}
	return nil
}
