// Crawl a website, depth-first, listing all unique paths found
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	foundPaths  []string
	startingUrl *url.URL
	timeout     = time.Duration(8 * time.Second)
)

func crawlUrl(path string) {
	// Create a temporary URL object for this request
	var targetUrl url.URL
	targetUrl.Scheme = startingUrl.Scheme
	targetUrl.Host = startingUrl.Host
	targetUrl.Path = path

	// Fetch the URL with a timeout and parse to goquery doc
	httpClient := http.Client{Timeout: timeout}
	response, err := httpClient.Get(targetUrl.String())
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	// Find all links and crawl if new path on same host
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		parsedUrl, err := url.Parse(href)
		if err != nil { // Err parsing URL. Ignore
			return
		}

		if urlIsInScope(parsedUrl) {
			foundPaths = append(foundPaths, parsedUrl.Path)
			log.Println("Found new path to crawl: " + parsedUrl.String())
			crawlUrl(parsedUrl.Path)
		}
	})
}

// Determine if path has already been found
// and if it points to the same host
func urlIsInScope(tempUrl *url.URL) bool {
	// Relative url, same host
	if tempUrl.Host != "" && tempUrl.Host != startingUrl.Host {
		return false // Link points to different host
	}

	if tempUrl.Path == "" {
		return false
	}

	// Already found?
	for _, existingPath := range foundPaths {
		if existingPath == tempUrl.Path {
			return false // Match
		}
	}
	return true // No match found
}

func main() {
	// Load command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Crawl a website, depth-first")
		fmt.Println("Usage: " + os.Args[0] + " <startingUrl>")
		fmt.Println("Example: " + os.Args[0] + " https://www.devdungeon.com")
		os.Exit(1)
	}
	foundPaths = make([]string, 0)

	// Parse starting URL
	startingUrl, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal("Error parsing starting URL. ", err)
	}
	log.Println("Crawling: " + startingUrl.String())

	crawlUrl(startingUrl.Path)

	for _, path := range foundPaths {
		fmt.Println(path)
	}
	log.Printf("Total unique paths crawled: %d\n", len(foundPaths))
}
