package service

import "github.com/Sonu875/goLearning/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerByID(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func NewCustomerService(repo domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}

}
