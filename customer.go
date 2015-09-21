package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Customer struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Address     string        `json:"address"`
	Address2    string        `json:"address2"`
	City        string        `json:"city"`
	Zip         string        `json:"zip"`
	Country     string        `json:"country"`
	Telephone   string        `json:"telephone"`
	Telephone2  string        `json:"telephone2"`
	DateCreated time.Time     `json:"date_created"`
}

func (customer *Customer) valid() bool {
	return len(customer.Id) > 0 && len(customer.Name) > 0 && len(customer.Email) > 0
}

func (customer *Customer) printConnectionDetails() {

	fmt.Println("Name: ", customer.Name)
	fmt.Println("Email: ", customer.Email)

}

func (customer *Customer) CreateNewCustomer(mConnection *MongoConnection) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	collection := session.DB("goinvoice").C("customerdata")
	err := collection.Insert(customer)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (customer *Customer) UpdateCustomer(mConnection *MongoConnection) bool {
	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("goinvoice").C("customerdata")
	err := c.UpdateId(customer.Id, customer)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (customer *Customer) DeleteCustomer(mConnection *MongoConnection, token string, customerID string) bool {
	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("goinvoice").C("customerdata")
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(customerID), "$and": []interface{}{bson.M{"user_id": bson.ObjectIdHex(token)}}})

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (customer *Customer) ListCustomers(mConnection *MongoConnection, token string) []Customer {
	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var customers Customers

	c := session.DB("goinvoice").C("customerdata")
	err := c.Find(bson.M{"user_id": bson.ObjectIdHex(token)}).All(&customers)
	if err != nil {
		log.Fatal(err)
	}

	return customers
}

func (customer *Customer) FindCustomerByID(mConnection *MongoConnection, token string, customerID string) *Customer {
	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var customerData *Customer

	c := session.DB("goinvoice").C("customerdata")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(customerID), "$and": []interface{}{bson.M{"user_id": bson.ObjectIdHex(token)}}}).One(&customerData)
	if err != nil {
		log.Fatal(err)
	}

	return customerData
}

type Customers []Customer
