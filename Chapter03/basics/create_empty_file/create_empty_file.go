package main

import (
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
