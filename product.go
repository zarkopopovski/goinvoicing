package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
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

func (product *Product) ListProducts(mConnection *MongoConnection, token string) []Product {
	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var products Products

	c := session.DB("goinvoice").C("productdata")
	err := c.Find(bson.M{"user_id": bson.ObjectIdHex(token)}).All(&products)
	if err != nil {
		log.Fatal(err)
	}

	return products
}

func (product *Product) FindProduct(mConnection *MongoConnection, token string, productID string) *Product {
	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	var productData *Product

	c := session.DB("goinvoice").C("productdata")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(productID), "$and": []interface{}{bson.M{"user_id": bson.ObjectIdHex(token)}}}).One(&productData)
	if err != nil {
		log.Fatal(err)
	}

	return productData
}

func (product *Product) FindProductsByIDs(mConnection *MongoConnection, token string, productsIDs []string) []Product {
	if mConnection.dbSession == nil {
		return nil
	}

	session := mConnection.dbSession.Copy()
	defer session.Close()

	objectsArray := make([]bson.ObjectId, len(productsIDs))

	for i := 0; i < len(productsIDs); i++ {
		objectsArray[i] = bson.ObjectIdHex(productsIDs[i])
	}

	var products Products

	c := session.DB("goinvoice").C("productdata")
	err := c.Find(
		bson.M{"user_id": bson.ObjectIdHex(token),
			"$and": []interface{}{
				bson.M{"_id": []interface{}{
					bson.M{"$in": objectsArray}}}}}).All(&products)
	if err != nil {
		log.Fatal(err)
	}

	return products
}

type Products []Product
