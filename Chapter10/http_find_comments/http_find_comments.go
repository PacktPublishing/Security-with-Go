// Search through a URL and find HTML comments
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	// Load command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Search for HTML comments in a URL")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]

	// Fetch the URL and get response
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Look for HTML comments using a regular expression
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	matches := re.FindAllString(string(body), -1)
	if matches == nil {
		// Clean exit if no matches found
		fmt.Println("No HTML comments found.")
		os.Exit(0)
	}

	// Print all HTML comments found
	for _, match := range matches {

		//cleanedMatch := match[4 : len(match)-1]
		fmt.Println(match)
	}
}
