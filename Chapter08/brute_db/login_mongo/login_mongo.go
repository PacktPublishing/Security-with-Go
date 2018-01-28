package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

func main() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"example.com"},
		Timeout:  30 * time.Second,
		Database: "dbname",
		Username: "username",
		Password: "password",
	}
	_, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Println("Error connecting. ", err)
	} else { // Success
		log.Println("Successfull connection.")
	}

}
