package domain

import appError "github.com/Sonu875/goLearning/Errors"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepo interface {
	FindAll(status string) ([]Customer, *appError.AppError)
	GetCustomerByID(id string) (*Customer, *appError.AppError)
}
