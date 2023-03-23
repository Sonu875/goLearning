package service

import (
	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/domain"
	"github.com/Sonu875/goLearning/dto"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *appError.AppError)
	GetCustomerByID(id string) (*dto.CustomerResponse, *appError.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *appError.AppError) {
	return s.repo.FindAll(status)

}

func (s DefaultCustomerService) GetCustomerByID(id string) (*dto.CustomerResponse, *appError.AppError) {

	c, err := s.repo.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}

}
