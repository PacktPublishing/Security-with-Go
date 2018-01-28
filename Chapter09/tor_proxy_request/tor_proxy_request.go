package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// The Tor proxy server must already be running and listening
func main() {
	targetUrl := "https://check.torproject.org"
	torProxy := "socks5://localhost:9050" // 9150 w/ Tor Browser

	// Parse Tor proxy URL string to a URL type
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ". ", err)
	}

	// Set up a custom HTTP transport for the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{
		Transport: torTransport,
		Timeout:   time.Second * 5,
	}

	// Make request
	response, err := client.Get(targetUrl)
	if err != nil {
		log.Fatal("Error making GET request. ", err)
	}
	defer response.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body of response. ", err)
	}
	log.Println(string(body))
}
