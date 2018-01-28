// Perform an HTTP HEAD request on a URL and print out headers
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load URL from command line arguments
	if len(os.Args) != 2 {
		fmt.Println(os.Args[0] + " - Perform an HTTP HEAD request to a URL")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]

	// Perform HTTP HEAD
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	// Print out each header key and value pair
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value[0])
	}
}
