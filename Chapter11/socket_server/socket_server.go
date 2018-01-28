package main

import (
	"net"
	"fmt"
	"log"
)

var protocol = "tcp" // tcp or udp
var listenAddress = "localhost:3000"

func main() {
	listener, err := net.Listen(protocol, listenAddress)
	if err != nil {
		log.Fatal("Error creating listener. ", err)
	}
	log.Printf("Now listening for connections.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection. ", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	incomingMessageBuffer := make([]byte, 4096)

	numBytesRead, err := conn.Read(incomingMessageBuffer)
	if err != nil {
		log.Print("Error reading from client. ", err)
	}

	fmt.Fprintf(conn, "Thank you. I processed %d bytes.\n", numBytesRead)
}
