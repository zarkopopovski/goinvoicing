package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name         string        `json:"user_name"`
	Email        string        `json:"email"`
	Ages         int           `json:"ages"`
	Sex          string        `json:"sex"`
	Password     string        `json:"password"`
	Date_Created time.Time     `json:"date_created"`
}

func (user *User) valid() bool {
	return len(user.Id) > 0 && len(user.Name) > 0 && len(user.Email) > 0
}

func (user *User) printConnectionDetails() {

	fmt.Println("Name: ", user.Name)
	fmt.Println("Ages: ", user.Ages)
	fmt.Println("Email: ", user.Email)

}

func (user *User) CreateNewUser(mConnection *MongoConnection) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	collection := session.DB("goinvoice").C("userdata")
	err := collection.Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (user *User) LoginWithCredentials(mConnection *MongoConnection, email string, password string) *User {

	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("goinvoice").C("userdata")
	err := c.Find(bson.M{"email": email, "$and": []interface{}{bson.M{"password": password}}}).One(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

type Users []User
