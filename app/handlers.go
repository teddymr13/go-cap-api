package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"User 1", "Jakarta", "12345"},
		{"User 2", "Surabaya", "67890"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
