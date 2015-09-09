package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"models"
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
