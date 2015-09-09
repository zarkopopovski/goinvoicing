package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID       bson.ObjectId `bson:"user_id" json:"user_id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Price        string        `json:"price"`
	Tax          string        `json:"tax"`
	Valid        bool          `json:"valid"`
	Date_Created time.Time     `json:"date_created"`
}

type Products []Product
