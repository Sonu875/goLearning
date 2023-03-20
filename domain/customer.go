package domain

import appError "github.com/Sonu875/goLearning/Errors"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepo interface {
	FindAll() ([]Customer, *appError.AppError)
	GetCustomerByID(id string) (*Customer, *appError.AppError)
}
