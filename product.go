package main

import (
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

func (product *Product) CreateNewProduct(mConnection *MongoConnection) bool {

	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	collection := session.DB("goinvoice").C("productdata")
	err := collection.Insert(product)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (product *Product) UpdateProduct(mConnection *MongoConnection) bool {
	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("goinvoice").C("productdata")
	err := c.UpdateId(product.Id, product)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (product *Product) DeleteProduct(mConnection *MongoConnection, token string, productID string) bool {
	if mConnection.dbSession == nil {
		return false
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	c := session.DB("goinvoice").C("productdata")
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(productID), "$and": []interface{}{bson.M{"user_id": bson.ObjectIdHex(token)}}})

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

type Products []Product
