package main

import (
"fmt"
"log"
"net/http"
"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Serve a directory via HTTP

URL should include protocol IP or hostname and port separated by colon.

Usage:
  ` + os.Args[0] + ` <listenUrl> <directory>

Example:
  ` + os.Args[0] + ` localhost:8080 .
  ` + os.Args[0] + ` 0.0.0.0:9999 /home/nanodano
`)
}

func checkArgs() (string, string) {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1], os.Args[2]
}

func main() {
	listenUrl, directoryPath := checkArgs()
	err := http.ListenAndServe(listenUrl, http.FileServer(http.Dir(directoryPath)))
	if err != nil {
		log.Fatal("Error running server. ", err)
	}
}

