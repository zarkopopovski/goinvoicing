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

func (mConnection *MongoConnection) CreateNewUser(u *User) bool {
	return u.CreateNewUser(mConnection)
}

func (mConnection *MongoConnection) LoginWithCredentials(email string, password string) *User {
	user := &User{}
	return user.LoginWithCredentials(mConnection, email, password)
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
