package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	Session *mgo.Session
)

func init() {
	var err error
	Session, err = mgo.DialWithTimeout("mongodb://localhost:27017/admin", 10)
	if err != nil {
		log.Fatal("mongodb init fatal!")
	}
}

