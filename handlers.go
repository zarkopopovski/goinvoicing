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

/*
func (c *ApiConnection) TestRoute(w http.ResponseWriter, r *http.Request) {

	products := &Products{
		Product{Id: bson.NewObjectId(), UserID: bson.NewObjectId(), Name: "Test Product", Description: "Test Description", Price: "1000", Tax: "18", Valid: true},
		Product{Id: bson.NewObjectId(), UserID: bson.NewObjectId(), Name: "Test Product2", Description: "Test Description", Price: "1000", Tax: "18", Valid: true},
		Product{Id: bson.NewObjectId(), UserID: bson.NewObjectId(), Name: "Test Product3", Description: "Test Description", Price: "1000", Tax: "18", Valid: true}}

	invoce := &Invoice{Id: bson.NewObjectId(), UserID: bson.NewObjectId(), Products: *products, Description: "Test, Test, Test, Test, Test", Price: "41242412412412", Valid: true}

	c.dbConnection.SaveTestObject(invoce)

}
*/

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
