package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "Jakarta", "12345", "2022-01-01", "1"},
		{"2", "User2", "Surabaya", "67890", "2022-01-01", "1"},
		{"3", "User3", "Jakarta", "12345", "2022-01-01", "1"},
		{"4", "User4", "Surabaya", "67890", "2022-01-01", "1"},
	}

	return CustomerRepositoryStub{Customer: customers}
}
