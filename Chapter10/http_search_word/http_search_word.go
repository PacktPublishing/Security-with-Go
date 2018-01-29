// Perform an HTTP request to load a page and search for a string
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Load command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Search for a keyword in the contents of a URL")
		fmt.Println("Usage: " + os.Args[0] + " <url> <keyword>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com NanoDano")
		os.Exit(1)
	}
	url := os.Args[1]
	needle := os.Args[2]

	// Create a custom http client to override default settings. Optional step.
	// Use http.Get() instead of client.Get() to use default client.
	client := &http.Client{
		Timeout: 30 * time.Second, // Default is forever!
		// CheckRedirect - Policy for following HTTP redirects
		// Jar - Cookie jar holding cookies
		// Transport - Change default method for making request
	}

	response, err := client.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Search for string
	if strings.Contains(string(body), needle) {
		fmt.Println("Match found for " + needle + " in URL " + url)
	} else {
		fmt.Println("No match found for " + needle + " in URL " + url)
	}
}
