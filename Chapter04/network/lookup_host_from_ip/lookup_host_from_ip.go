package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No IP address argument provided.")
	}
	arg := os.Args[1]

	// Parse the IP for validation
	ip := net.ParseIP(arg)
	if ip == nil {
		log.Fatal("Valid IP not detected. Value provided: " + arg)
	}

	fmt.Println("Looking up hostnames for IP address: " + arg)
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostnames := range hostnames {
		fmt.Println(hostnames)
	}
}
