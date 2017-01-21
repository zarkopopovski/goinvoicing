package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Invoice struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID      bson.ObjectId `json:"user_id" bson:"user_id"`
	Customer    Customer      `json:"customer" bson:"customer"`
	Products    Products      `json:"products" bson:"products"`
	Description string        `json:"description" bson:"description"`
	Price       string        `json:"price" bson:"price"`
	Valid       bool          `json:"valid" bson:"valid"`
	DateCreated time.Time     `json:"date_created" bson:"date_created"`
}

func (invoice *Invoice) CreateInvoice(mConnection *MongoConnection) bool {

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
