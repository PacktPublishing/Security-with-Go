package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	proxyUrlString := "http://<proxyIp>:<proxyPort>"
	proxyUrl, err := url.Parse(proxyUrlString)
	if err != nil {
		log.Fatal("Error parsing URL. ", err)
	}

	// Set up a custom HTTP transport for client
	customTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	httpClient := &http.Client{
		Transport: customTransport,
		Timeout:   time.Second * 5,
	}

	// Make request
	response, err := httpClient.Get("http://www.example.com")
	if err != nil {
		log.Fatal("Error making GET request. ", err)
	}
	defer response.Body.Close()

	// Read and print response from server
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body of response. ", err)
	}
	log.Println(string(body))

}
