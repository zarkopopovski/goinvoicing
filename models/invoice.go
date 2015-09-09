package models

import (
	"fmt"
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

type Invoices []Invoice
