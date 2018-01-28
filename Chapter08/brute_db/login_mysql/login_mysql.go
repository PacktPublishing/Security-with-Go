package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	_, err := sql.Open("mysql", "user:password@hostname/dbname?charset=utf8")
	if err != nil {
		log.Println("Error connecting. ", err)
	} else { // Success
		log.Println("Successfull connection.")
	}
}
