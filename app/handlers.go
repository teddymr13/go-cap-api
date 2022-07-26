package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
	Phone   string `json:"telepon" xml:"phone"`
}

var customers []Customer = []Customer{
	{1, "User1", "Jakarta", "123456", "0812732323"},
	{2, "User2", "Surabaya", "67890", "0832231231"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint")
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Fprintln(w, vars["customer_id"])

	customerId := vars["customer_id"]
	id, _ := strconv.Atoi(customerId)

	var cust Customer
	for _, data := range customers {
		if data.ID == id {
			cust = data
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}
