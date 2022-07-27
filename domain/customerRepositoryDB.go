package domain

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:postgres@localhost/banking?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, error) {
	query := "select * from customers where customer_id = $1"

	row := d.client.QueryRow(query, customerID)

	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		log.Println("error scanning customer data ", err.Error())
		return nil, err
	}

	return &c, nil
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	query := "select * from customers"

	rows, err := d.client.Query(query)
	if err != nil {
		log.Println("error query data to customer table ", err.Error())
		return nil, err
	}

	var customers []Customer
	for rows.Next() {

		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer data ", err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}
