package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint\n")

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {

	// * get route variable
	vars := mux.Vars(r)

	customerID := vars["customer_id"]

	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	// * return customer data
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// * decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// * generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// * save data to array
// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer successfully created")
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) {

// 	// * get route variable
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// * convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}

// 	// * searching customer data
// 	var cust Customer

// 	for customerIndex, data := range customers {
// 		if data.ID == id {
// 			// * save temp data for validation
// 			cust = data

// 			// * decode request body
// 			var newCust Customer
// 			json.NewDecoder(r.Body).Decode(&newCust)

// 			// * do update
// 			customers[customerIndex].Name = newCust.Name
// 			customers[customerIndex].City = newCust.City
// 			customers[customerIndex].ZipCode = newCust.ZipCode

// 			// fmt.Println(customers)

// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintln(w, "customer data updated")
// 			return
// 		}
// 	}

// 	// ! error if data not find
// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}

// }

// func getNextID() int {
// 	lastIndex := len(customers) - 1
// 	lastCustomer := customers[lastIndex]
// 	// cust := customers[len(customers)-1]

// 	return lastCustomer.ID + 1
// }
