package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load command line arguments
	if len(os.Args) != 2 {
		fmt.Println("List all JavaScript files in a webpage")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	url := os.Args[1]

	// Fetch the URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and list all external scripts in page
	fmt.Println("Scripts found in", url)
	fmt.Println("==========================")
	doc.Find("script").Each(func(i int, script *goquery.Selection) {

		// By looking only at the script src we are limiting
		// the search to only externally loaded JavaScript files.
		// External files might be hosted on the same domain
		// or hosted remotely
		src, exists := script.Attr("src")
		if exists {
			fmt.Println(src)
		}

		// script.Text() will contain the raw script text
		// if the JavaScript code is written directly in the
		// HTML source instead of loaded from a separate file
	})
}
