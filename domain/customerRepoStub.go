package domain

type CustomerRepoStub struct {
	customers []Customer
}

func (s CustomerRepoStub) FindAll() ([]Customer, error) {

	return s.customers, nil
}

func NewCustomerRepoStub() CustomerRepoStub {
	customers := []Customer{
		{Id: "10001", Name: "Sonu", City: "Derabassi", Zipcode: "140507", DateOfBirth: "26/08/1996", Status: "1"},
		{Id: "10002", Name: "Sonu", City: "Derabassi", Zipcode: "140507", DateOfBirth: "26/08/1996", Status: "1"},
	}
	return CustomerRepoStub{customers: customers}
}
