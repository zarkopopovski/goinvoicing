package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

type ProductsHandlers struct {
	dbConnection *MongoConnection
}

func (pHandlers *ProductsHandlers) NewProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	tax := r.FormValue("tax")

	validProduct, _ := strconv.ParseBool(r.FormValue("valid"))

	product := &Product{
		UserID:       bson.ObjectIdHex(token),
		Name:         name,
		Description:  description,
		Price:        price,
		Tax:          tax,
		Valid:        validProduct,
		Date_Created: time.Now()}

	result := pHandlers.dbConnection.CreateNewProduct(product)

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

func (pHandlers *ProductsHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	tax := r.FormValue("tax")

	validProduct, _ := strconv.ParseBool(r.FormValue("valid"))

	product := &Product{
		Id:          bson.ObjectIdHex(productID),
		UserID:      bson.ObjectIdHex(token),
		Name:        name,
		Description: description,
		Price:       price,
		Tax:         tax,
		Valid:       validProduct}

	result := pHandlers.dbConnection.UpdateProduct(product)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error updating product"}); err != nil {
		panic(err)
	}
}

func (pHandlers *ProductsHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	result := pHandlers.dbConnection.DeleteProduct(token, productID)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Error removing product"}); err != nil {
		panic(err)
	}
}

func (pHandlers *ProductsHandlers) ListProducts(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	products := pHandlers.dbConnection.ListAllProducts(token)

	if len(products) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(products); err != nil {
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
