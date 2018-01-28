package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Make basic HTTP GET request
	response, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	// Read body from response
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Printf("%s\n", body)
}
