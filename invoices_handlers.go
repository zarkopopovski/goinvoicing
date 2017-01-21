package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

func (c *ApiConnection) NewInvoice(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	customerID := r.FormValue("customer_id")
	description := r.FormValue("description")
	price := r.FormValue("price")

	validInvoice, _ := strconv.ParseBool(r.FormValue("valid"))

	products := []string{"5601084ea3dca2f03e9fd5c1", "560110fba3dca2f03e9fd5c4"}

	var customer = c.dbConnection.FindExistingCustomers(token, customerID)
	var productsData = c.dbConnection.FindExistingProducts(token, products)

	invoice := &Invoice{
		UserID:      bson.ObjectIdHex(token),
		Customer:    *customer,
		Products:    productsData,
		Description: description,
		Price:       price,
		Valid:       validInvoice,
		DateCreated: time.Now()}

	result := c.dbConnection.CreateNewInvoice(invoice)

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

func (c *ApiConnection) UpdateInvoice(w http.ResponseWriter, r *http.Request) {

}

func (c *ApiConnection) DeleteInvoice(w http.ResponseWriter, r *http.Request) {

}

func (c *ApiConnection) ListInvoices(w http.ResponseWriter, r *http.Request) {

}
