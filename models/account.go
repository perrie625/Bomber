package models

import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

var (
	collection = "Accounts"
	database = "users"
)


type account struct {
	// 因为json不可访问struct的
	Id bson.ObjectId `bson:"_id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Name string `bson:"Name"`
}

func (a *account)SetPassword(password string) error {
	errResult := errors.New("Set Password Failed")
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return errResult
	}
	a.Password = string(hashedPassword)
	// update data in db
	return nil
}


func NewAccount(username string, password string)(*account, error) {
	result := &account{
		Username: username,
	}
	err := result.SetPassword(password)
	if err != nil {
		return nil, errors.New("New Account Failed")
	}
	return result, nil
}


//func (a *account)Create() error{
//
//}


//func CreateAccount(name string, password string)(*account, error) {
//
//	passwordBytes := bytes.
//	hasedPassword := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
//
//	session := db.Session.Clone()
//	defer session.Close()
//
//	var account account
//	c := session.DB(database).C(collection)
//	err := c.Find(bson.M{"name": name}).One(&account)
//	if account.Username == ""{
//		err = error("no account")
//	}
//	return &account, err
//}
//
//
//
//func FindAccount(name string)(*account, error) {
//	session := db.Session.Clone()
//	defer session.Close()
//
//	var account account
//	c := session.DB(database).C(collection)
//	err := c.Find(bson.M{"name": name}).One(&account)
//	if account.Username == ""{
//		err = error("no account")
//	}
//	return &account, err
//}

