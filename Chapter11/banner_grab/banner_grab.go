package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "127.0.0.1"

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := 0; port <= 1024; port++ {
		go grabBanner(ipToScan, port, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func grabBanner(ip string, port int, doneChannel chan bool) {
	connection, err := net.DialTimeout(
		"tcp",
		ip+":"+strconv.Itoa(port),
		time.Second*10,
	)
	if err != nil {
		doneChannel <- true
		return
	}

	// See if server offers anything to read
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5)) // Set timeout
	numBytesRead, err := connection.Read(buffer)
	if err != nil {
		doneChannel <- true
		return
	}
	log.Printf("Banner from port %d\n%s\n", port, buffer[0:numBytesRead])

	doneChannel <- true
}
