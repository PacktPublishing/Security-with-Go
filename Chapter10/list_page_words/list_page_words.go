package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Load command line arguments
	if len(os.Args) != 2 {
		fmt.Println("List all words by frequency from a web page")
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

	// Find and list all headings h1-h6
	wordCountMap := make(map[string]int)
	doc.Find("p").Each(func(i int, body *goquery.Selection) {
		fmt.Println(body.Text())
		words := strings.Split(body.Text(), " ")
		for _, word := range words {
			trimmedWord := strings.Trim(word, " \t\n\r,.?!")
			if trimmedWord == "" {
				continue
			}
			wordCountMap[strings.ToLower(trimmedWord)]++

		}
	})

	// Print all words along with the number of times the word was seen
	for word, count := range wordCountMap {
		fmt.Printf("%d | %s\n", count, word)
	}

}
