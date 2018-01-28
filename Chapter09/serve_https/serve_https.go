package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested: "+request.URL.Path)
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServeTLS(
		"localhost:8181",
		"cert.pem",
		"privateKey.pem",
		nil,
	)
	if err != nil {
		log.Fatal("Error creating server. ", err)
	}
}
