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
		fmt.Println("List all headings (h1-h6) in a web page")
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

	// Print title before headings
	title := doc.Find("title").Text()
	fmt.Printf("== Title ==\n%s\n", title)

	// Find and list all headings h1-h6
	headingTags := [6]string{"h1", "h2", "h3", "h4", "h5", "h6"}
	for _, headingTag := range headingTags {
		fmt.Printf("== %s ==\n", headingTag)
		doc.Find(headingTag).Each(func(i int, heading *goquery.Selection) {
			fmt.Println(" * " + heading.Text())
		})
	}

}
