package main

import (
	"bytes"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	log.Printf("Received connection from %s.\n", conn.RemoteAddr())
	buff := make([]byte, 1024)
	nbytes, err := conn.Read(buff)
	if err != nil {
		log.Println("Error reading from connection. ", err)
	}
	// Always reply with a fake auth failed message
	conn.Write([]byte("Authentication failed."))
	trimmedOutput := bytes.TrimRight(buff, "\x00")
	log.Printf("Read %d bytes from %s.\n%s\n",
		nbytes, conn.RemoteAddr(), trimmedOutput)
	conn.Close()
}

func main() {
	portNumber := "9001" // or os.Args[1]
	ln, err := net.Listen("tcp", "localhost:"+portNumber)
	if err != nil {
		log.Fatalf("Error listening on port %s.\n%s\n",
			portNumber, err.Error())
	}
	log.Printf("Listening on port %s.\n", portNumber)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection.", err)
		}
		go handleConnection(conn)
	}
}
