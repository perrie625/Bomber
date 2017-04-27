package models

import (
	"Bomber/db"
	"gopkg.in/mgo.v2/bson"
)

var (
	collection = "Accounts"
	database = "users"
)


type Account struct {
	Name string `json:"name"`
	Password string `json:"password"`
}



func FindAccount(name string)(*Account, error) {
	session := db.Session.Clone()
	defer session.Close()

	var account Account
	c := session.DB(database).C(collection)
	err := c.Find(bson.M{"name": name}).One(&account)
	if account.Name == ""{
		err = error("no account")
	}
	return &account, err
}

