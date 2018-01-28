package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal("Error creating file.")
	}
	defer file.Close()
	// It is important to defer after checking the errors.
	// You can't call Close() on a nil object
	// if the open failed.

	// Some other actions

	// file.Close() will be called before final exit
}
