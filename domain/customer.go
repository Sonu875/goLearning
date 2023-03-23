package domain

import (
	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/dto"
)

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

func (c Customer) ToDto() dto.CustomerResponse {
	status := ""
	if c.Status == "0" {
		status = "inactive"
	} else {
		status = "active"
	}

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      status,
	}

}
