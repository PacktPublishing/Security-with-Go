package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No hostname argument provided.")
	}
	arg := os.Args[1]

	fmt.Println("Looking up IP addresses for hostname: " + arg)

	ips, err := net.LookupHost(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
