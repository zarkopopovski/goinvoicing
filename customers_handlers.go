package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type CustomersHandlers struct {
	dbConnection *MongoConnection
}

func (cHandlers *CustomersHandlers) NewCustomer(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")
	address2 := r.FormValue("address2")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	telephone := r.FormValue("telephone")
	telephone2 := r.FormValue("telephone2")

	customer := &Customer{
		UserID:      bson.ObjectIdHex(token),
		Name:        name,
		Email:       email,
		Address:     address,
		Address2:    address2,
		City:        city,
		Zip:         zip,
		Country:     country,
		Telephone:   telephone,
		Telephone2:  telephone2,
		DateCreated: time.Now()}

	result := cHandlers.dbConnection.CreateNewCustomer(customer)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (cHandlers *CustomersHandlers) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	customerID := r.FormValue("customer_id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")
	address2 := r.FormValue("address2")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	telephone := r.FormValue("telephone")
	telephone2 := r.FormValue("telephone2")

	customer := &Customer{
		Id:          bson.ObjectIdHex(customerID),
		UserID:      bson.ObjectIdHex(token),
		Name:        name,
		Email:       email,
		Address:     address,
		Address2:    address2,
		City:        city,
		Zip:         zip,
		Country:     country,
		Telephone:   telephone,
		Telephone2:  telephone2,
		DateCreated: time.Now()}

	result := cHandlers.dbConnection.UpdateCustomer(customer)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (cHandlers *CustomersHandlers) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	customerID := r.FormValue("customer_id")

	result := cHandlers.dbConnection.DeleteCustomer(token, customerID)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error removing customer"}); err != nil {
		panic(err)
	}
}

func (cHandlers *CustomersHandlers) ListCustomers(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	customers := cHandlers.dbConnection.ListAllCustomers(token)

	if len(customers) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
