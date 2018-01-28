package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No domain name argument provided")
	}
	arg := os.Args[1]

	fmt.Println("Looking up MX records for " + arg)

	mxRecords, err := net.LookupMX(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, mxRecord := range mxRecords {
		fmt.Printf("Host: %s\tPreference: %d\n", mxRecord.Host, mxRecord.Pref)
	}
}
