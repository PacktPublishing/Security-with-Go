package main

import (
	"log"
	"time"
)

func countDown() {
	for i := 5; i >= 0; i-- {
		log.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Kick off a thread
	go countDown()

	// Since functions are first-class
	// you can write an anonymous function
	// for a go routine
	go func() {
		time.Sleep(time.Second * 2)
		log.Println("Delayed greetings!")
	}()

	// Use channels to signal when complete
	// Or in this case just wait
	time.Sleep(time.Second * 4)
}
