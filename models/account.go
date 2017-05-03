package models

import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"Bomber/db"
	"gopkg.in/mgo.v2"
	"Bomber/protodata"
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

func (a *account)Create()  error {
//	新建账号到数据库
	session := db.Session.Clone()
	defer session.Close()

	c := session.DB(database).C(collection)
	err := c.Insert(a)
	return err
}


func CreateAccount(l *protodata.LoginRequest)(*account, error) {
	result := &account{
		Id: bson.NewObjectId(),
		Username: l.Username,
		Name: "dafultName",
	}
	err := result.SetPassword(l.Password)
	if err != nil {
		return nil, errors.New("Create Account Failed")
	}
	// todo: error handle
	result.Create()
	return result, nil
}





func ensureIndex(){
	session := db.Session.Clone()
	defer session.Close()

	c := session.DB(database).C(collection)
	index := mgo.Index{
		Key: []string{"username"},
		Unique: true,
		DropDups: true,
		Background:true,
		Sparse:true,
	}
	err := c.EnsureIndex(index)
	if err != nil{
		panic(err)
	}
}

func init() {
	ensureIndex()
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

