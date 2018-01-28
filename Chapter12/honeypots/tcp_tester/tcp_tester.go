package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func checkArgs() string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <targetAddress>")
		fmt.Println("Example: " + os.Args[0] + " localhost:9001")
		os.Exit(0)
	}
	return os.Args[1]
}

func main() {
	var err error
	targetAddress := checkArgs()
	conn, err := net.Dial("tcp", targetAddress)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)

	_, err = os.Stdin.Read(buf)
	trimmedInput := bytes.TrimRight(buf, "\x00")
	log.Printf("%s\n", trimmedInput)

	_, writeErr := conn.Write(trimmedInput)
	if writeErr != nil {
		log.Fatal("Error sending data to remote host. ", writeErr)
	}

	_, readErr := conn.Read(buf)
	if readErr != nil {
		log.Fatal("Error when reading from remote host. ", readErr)
	}
	trimmedOutput := bytes.TrimRight(buf, "\x00")
	log.Printf("%s\n", trimmedOutput)
}
