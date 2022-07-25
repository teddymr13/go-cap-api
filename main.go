package main

import (
	"fmt"
	"net/http"
)

func main() {

	// * defining routes
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Celerates!")
	})

	// * starting the server
	http.ListenAndServe(":8080", nil)
}
