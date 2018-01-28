package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgres://user:password@hostname/dbname?sslmode=verify-full"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error connecting. ", err)
	} else { // Success
		log.Println("Successfull connection.")
	}
}
