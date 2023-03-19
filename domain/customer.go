package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepo interface {
	FindAll() ([]Customer, error)
}
