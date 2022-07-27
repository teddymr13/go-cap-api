package domain

type Customer struct {
	ID          int    `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipcode"`
	Phone       string `json:"telepon" xml:"phone"`
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
