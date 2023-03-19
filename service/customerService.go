package service

import "github.com/Sonu875/goLearning/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}

}
