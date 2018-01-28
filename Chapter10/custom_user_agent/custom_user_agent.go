// Change HTTP user agent
package main

import (
	"log"
	"net/http"
)

func main() {
	// Create the request for use later
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
	}

	// Override the user agent
	request.Header.Set("User-Agent", "_Custom User Agent_")

	// Perform the request, ignore response.
	_, err = client.Do(request)
	if err != nil {
		log.Fatal("Error making request. ", err)
	}
}
