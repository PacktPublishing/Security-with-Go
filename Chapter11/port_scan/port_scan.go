package main

import (
	"strconv"
	"log"
	"net"
	"time"
)

var ipToScan = "127.0.0.1"
var minPort = 0
var maxPort = 1024

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= maxPort ; port++ {
		go testTcpConnection(ipToScan, port, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<- doneChannel
		activeThreads--
	}
}

func testTcpConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip + ":" + strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Port %d: Open\n", port)
	}
	doneChannel <- true
}