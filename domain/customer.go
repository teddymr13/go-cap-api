package domain

import (
	"capi/dto"
	"capi/errs"
)

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	//  status -> "1", "0", ""
	FindAll(string) ([]Customer, *errs.AppErr)
	FindByID(string) (*Customer, *errs.AppErr)
}

func (c Customer) convertStatusName() string {
	statusName := "active"
	if c.Status == "0" {
		statusName = "inactive"
	}
	return statusName
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		City:        c.City,
		ZipCode:     c.ZipCode,
		Status:      c.convertStatusName(),
	}
}
