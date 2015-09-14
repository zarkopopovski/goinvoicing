package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ApiConnection struct {
	dbConnection *MongoConnection
}

func CreateApiConnection() *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
	}
	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func (c *ApiConnection) SignIn(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	ages := r.FormValue("ages")
	sex := r.FormValue("sex")
	password := r.FormValue("password")

	userAges, _ := strconv.Atoi(ages)

	user := &User{Name: name, Email: email, Ages: userAges, Sex: sex, Password: password, Date_Created: time.Now()}

	result := c.dbConnection.CreateNewUser(user)

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

func (c *ApiConnection) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := c.dbConnection.LoginWithCredentials(email, password)
	userID := fmt.Sprintf("%x", string(user.Id))
	log.Println(userID)

	if user != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		userMap := make(map[string]string)
		userMap["token"] = userID

		if err := json.NewEncoder(w).Encode(userMap); err != nil {
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

func (c *ApiConnection) NewProduct(w http.ResponseWriter, r *http.Request) {
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

	result := c.dbConnection.CreateNewProduct(product)

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

func (c *ApiConnection) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	result := c.dbConnection.UpdateProduct(product)

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

func (c *ApiConnection) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	result := c.dbConnection.DeleteProduct(token, productID)

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

func (c *ApiConnection) ListProducts(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	products := c.dbConnection.ListAllProducts(token)

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

func (c *ApiConnection) NewCustomer(w http.ResponseWriter, r *http.Request) {
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

	result := c.dbConnection.CreateNewCustomer(customer)

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

func (c *ApiConnection) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
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

	result := c.dbConnection.UpdateCustomer(customer)

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

func (c *ApiConnection) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	customerID := r.FormValue("customer_id")

	result := c.dbConnection.DeleteCustomer(token, customerID)

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

func (c *ApiConnection) ListCustomers(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	customers := c.dbConnection.ListAllCustomers(token)

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
