package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// Read file to byte slice
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}
