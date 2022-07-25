package app

import (
	"log"
	"net/http"
)

func Start() {

	// * create multiplexer
	mux := http.NewServeMux()

	// * defining routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// * starting the server
	log.Fatal(http.ListenAndServe(":8080", mux))
}
