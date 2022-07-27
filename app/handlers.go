package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id"`
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// 	Phone   string `json:"telepon" xml:"phone"`
// }

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "123456", "0812732323"},
// 	{2, "User2", "Surabaya", "67890", "0832231231"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint")

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		panic(err.Error())
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// func getCustomers(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	// fmt.Fprintln(w, vars["customer_id"])

// 	// convert int to string
// 	customerId := vars["customer_id"]
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}

// 	//searching customer data
// 	var cust Customer
// 	for _, data := range customers {
// 		if data.ID == id {
// 			cust = data
// 		}
// 	}

// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}

// 	//return customer data
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cust)
// }

// func updateCustomers(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["customer_id"])
// 	for index, data := range customers {
// 		if data.ID == id {
// 			customers = append(customers[:index], customers[index+1:]...)
// 			var updateCustomers Customer

// 			json.NewDecoder(r.Body).Decode(&updateCustomers)
// 			customers = append(customers, updateCustomers)
// 			json.NewEncoder(w).Encode(updateCustomers)
// 			fmt.Println("update successfully")
// 			return
// 		}
// 	}
// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["customer_id"])
// 	for index, data := range customers {
// 		if data.ID == id {
// 			customers = append(customers[:index], customers[index+1:]...)
// 		}
// 	}
// 	fmt.Fprint(w, "Remove Succesfully")

// }

// func addCustomers(w http.ResponseWriter, r *http.Request) {
// 	// decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// save data to array
// 	customers = append(customers, cust)
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprint(w, "customer successfuly created")
// }

// func getNextID() int {
// 	cust := customers[len(customers)-1]
// 	return cust.ID + 1
// }
