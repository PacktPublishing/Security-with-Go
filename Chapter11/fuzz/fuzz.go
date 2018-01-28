package main

import (
	"crypto/rand"
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "www.devdungeon.com"
var port = 80
var maxFuzzBytes = 1024

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for fuzzSize := 1; fuzzSize <= maxFuzzBytes; fuzzSize = fuzzSize * 2 {
		go fuzz(ipToScan, port, fuzzSize, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func fuzz(ip string, port int, fuzzSize int, doneChannel chan bool) {
	log.Printf("Fuzzing %d.\n", fuzzSize)

	conn, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		log.Printf("Fuzz of %d attempted. Could not connect to server. %s\n", fuzzSize, err)
		doneChannel <- true
		return
	}

	// Write random bytes to server
	randomBytes := make([]byte, fuzzSize)
	rand.Read(randomBytes)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	numBytesWritten, err := conn.Write(randomBytes)
	if err != nil { // Error writing
		log.Printf(
			"Fuzz of %d attempted. Could not write to server. %s\n",
			fuzzSize,
			err,
		)
		doneChannel <- true
		return
	}
	if numBytesWritten != fuzzSize {
		log.Printf("Unable to write the full %d bytes.\n", fuzzSize)
	}
	log.Printf("Sent %d bytes:\n%s\n\n", numBytesWritten, randomBytes)

	// Read up to 4k back
	readBuffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	numBytesRead, err := conn.Read(readBuffer)
	if err != nil { // Error reading
		log.Printf(
			"Fuzz of %d attempted. Could not read from server. %s\n",
			fuzzSize,
			err,
		)
		doneChannel <- true
		return
	}

	log.Printf(
		"Sent %d bytes to server. Read %d bytes back:\n",
		fuzzSize,
		numBytesRead,
	)
	log.Printf(
		"Data:\n%s\n\n",
		readBuffer[0:numBytesRead],
	)

	doneChannel <- true
}
