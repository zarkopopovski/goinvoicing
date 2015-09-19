package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Invoice struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID       bson.ObjectId `bson:"user_id" json:"user_id"`
	Customer     Customer      `json:"customer"`
	Products     Products      `json:"products"`
	Description  string        `json:"description"`
	Price        string        `json:"price"`
	Valid        bool          `json:"valid"`
	Date_Created time.Time     `json:"date_created"`
}

func (invoice *Invoice) CreateNewInvoice(mConnection *MongoConnection) bool {

	return false

}

func (invoice *Invoice) UpdateExistingInvoice(mConnection *MongoConnection) bool {

	return false

}

func (invoice *Invoice) DeleteExistingInvoice(mConnection *MongoConnection) bool {

	return false

}

func (invoice *Invoice) ListExistingInvoices(mConnection *MongoConnection) []Invoice {

	invoices := Invoices{}

	return invoices

}

type Invoices []Invoice
