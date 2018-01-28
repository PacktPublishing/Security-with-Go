package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "https://www.example.com"

func main() {
	// Create the HTTP request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request. ", err)
	}

	// Set cookie
	request.Header.Set("Cookie", "session_id=<SESSION_TOKEN>")

	// Create the HTTP client, make request and print response
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	data, err := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", data)
}
