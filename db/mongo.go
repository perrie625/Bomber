package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"Bomber/tools"
)

var (
	Session *mgo.Session
)

func init() {
	var err error
	Session, err = mgo.DialWithTimeout(tools.ServerConfig.DBUrl, 10)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("mongodb init fatal!")
	}
}

