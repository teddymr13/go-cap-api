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
	mux.HandleFunc("/customers/{customer_id}", getCustomers).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
