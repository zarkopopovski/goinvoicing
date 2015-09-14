package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

type MongoConnection struct {
	dbSession *mgo.Session
}

func OpenConnectionSession() (mongoConnection *MongoConnection) {
	mongoConnection = new(MongoConnection)
	mongoConnection.createNewDBConnection()

	return
}

func (mConnection *MongoConnection) createNewDBConnection() (err error) {
	mConnection.dbSession, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	mConnection.dbSession.SetMode(mgo.Monotonic, true)

	return
}

func (mConnection *MongoConnection) CreateNewUser(user *User) bool {
	return user.CreateNewUser(mConnection)
}

func (mConnection *MongoConnection) LoginWithCredentials(email string, password string) *User {
	user := &User{}
	return user.LoginWithCredentials(mConnection, email, password)
}

func (mConnection *MongoConnection) CreateNewProduct(product *Product) bool {
	return product.CreateNewProduct(mConnection)
}

func (mConnection *MongoConnection) UpdateProduct(product *Product) bool {
	return product.UpdateProduct(mConnection)
}

func (mConnection *MongoConnection) DeleteProduct(token string, productID string) bool {
	product := &Product{}
	return product.DeleteProduct(mConnection, token, productID)
}

func (mConnection *MongoConnection) ListAllProducts(token string) []Product {
	product := &Product{}
	return product.ListProducts(mConnection, token)
}

func (mConnection *MongoConnection) CreateNewCustomer(customer *Customer) bool {
	return customer.CreateNewCustomer(mConnection)
}

func (mConnection *MongoConnection) UpdateCustomer(customer *Customer) bool {
	return customer.UpdateCustomer(mConnection)
}

func (mConnection *MongoConnection) DeleteCustomer(token string, customerID string) bool {
	customer := &Customer{}
	return customer.DeleteCustomer(mConnection, token, customerID)
}

func (mConnection *MongoConnection) ListAllCustomers(token string) []Customer {
	customer := &Customer{}
	return customer.ListCustomers(mConnection, token)
}

func (mConnection *MongoConnection) SaveTestObject(testInvoice *Invoice) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	collection := session.DB("goitest").C("invoice")
	err := collection.Insert(testInvoice)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}
