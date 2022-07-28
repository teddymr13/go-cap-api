package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			"1", "User1", "Jakarta", "181818", "2022-01-08", "1",
		},
		{
			"2", "User2", "Bandung", "989898", "2022-01-07", "1",
		},
		{
			"3", "User3", "Jogja", "00556", "2022-01-06", "1",
		},
		{
			"4", "User4", "Bali", "66999", "2022-01-05", "1",
		},
	}
	return CustomerRepositoryStub{Customer: customers}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {

	return s.Customer, nil
}
