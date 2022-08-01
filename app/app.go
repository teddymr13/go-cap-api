package app

import (
	"capi/domain"
	"capi/logger"
	"capi/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating application...", envKey))
		}
	}
}

func Start() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}

	sanityCheck()
	// * wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// * create ServeMux
	mux := mux.NewRouter()

	// * defining routes
	// mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// mux.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", updateCustomer).Methods(http.MethodPut)

	// * starting the server
	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("start server on %s:%s ...", serverAddr, serverPort))
	http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), mux)

}
