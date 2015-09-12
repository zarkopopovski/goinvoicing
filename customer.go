package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Customer struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID       bson.ObjectId `bson:"user_id" json:"user_id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Address      string        `json:"address"`
	Address2     string        `json:"address2"`
	City         string        `json:"city"`
	Zip          string        `json:"zip"`
	Country      string        `json:"country"`
	Telephone    string        `json:"telephone"`
	Telephone2   string        `json:"telephone2"`
	Date_Created time.Time     `json:"date_created"`
}

func (c *Customer) valid() bool {
	return len(c.Id) > 0 && len(c.Name) > 0 && len(c.Email) > 0
}

func (c *Customer) printConnectionDetails() {

	fmt.Println("Name: ", c.Name)
	fmt.Println("Email: ", c.Email)

}

type Customers []Customer
