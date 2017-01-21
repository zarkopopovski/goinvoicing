package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UsersHandlers struct {
	dbConnection *MongoConnection
}

func (uHandlers *UsersHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	ages := r.FormValue("ages")
	sex := r.FormValue("sex")
	password := r.FormValue("password")

	userAges, _ := strconv.Atoi(ages)

	user := &User{Name: name, Email: email, Ages: userAges, Sex: sex, Password: password, Date_Created: time.Now()}

	result := uHandlers.dbConnection.CreateNewUser(user)

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

func (uHandlers *UsersHandlers) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := uHandlers.dbConnection.LoginWithCredentials(email, password)
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
