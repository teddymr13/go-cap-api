package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//create ServerMux
	// mux := http.NewServeMux()
	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", addCustomers).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id}", updateCustomers).Methods(http.MethodPut)
	mux.HandleFunc("/customers/{customer_id}", deleteCustomer).Methods(http.MethodDelete)

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomers).Methods(http.MethodGet)
	// * starting the server
	http.ListenAndServe(":8080", mux)
}
