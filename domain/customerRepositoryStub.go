package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "100", Name: "ozan", City: "Jakrta", Zipcode: "123", DateOfBirth: "2000-03-22", Status: "1"},
		{Id: "101", Name: "fadhil", City: "Jakrta", Zipcode: "124", DateOfBirth: "1999-11-28", Status: "1"},
		{Id: "102", Name: "alfi", City: "Jakrta", Zipcode: "125", DateOfBirth: "1999-12-28", Status: "0"},
	}

	return CustomerRepositoryStub{customers: customers}
}