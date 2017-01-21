package main

import (
	"fmt"
	"net/http"
)

type ApiConnection struct {
	dbConnection *MongoConnection
	uHandlers    *UsersHandlers
	pHandlers    *ProductsHandlers
	cHandlers    *CustomersHandlers
}

func CreateApiConnection() *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
		uHandlers:    &UsersHandlers{},
		pHandlers:    &ProductsHandlers{},
		cHandlers:    &CustomersHandlers{},
	}
	API.uHandlers.dbConnection = API.dbConnection
	API.pHandlers.dbConnection = API.dbConnection
	API.cHandlers.dbConnection = API.dbConnection

	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
