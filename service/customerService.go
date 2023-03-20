package service

import (
	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *appError.AppError)
	GetCustomerByID(id string) (*domain.Customer, *appError.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *appError.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, *appError.AppError) {
	return s.repo.GetCustomerByID(id)
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}

}
