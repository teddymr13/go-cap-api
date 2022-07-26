package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// * create ServeMux
	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
