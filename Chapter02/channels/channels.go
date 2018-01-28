package main

import (
	"log"
	"time"
)

// Do some processing that takes a long time
// in a separate thread and signal when done
func process(doneChannel chan bool) {
	time.Sleep(time.Second * 3)
	doneChannel <- true
}

func main() {
	// Each channel can support one data type.
	// Can also use custom types
	var doneChannel chan bool

	// Channels are nil until initialized with make
	doneChannel = make(chan bool)

	// Kick off a lengthy process that will
	// signal when complete
	go process(doneChannel)

	// Get the first bool available in the channel
	// This is a blocking operation so execution
	// will not progress until value is received
	tempBool := <-doneChannel
	log.Println(tempBool)
	// or to simply ignore the value but still wait
	// <-doneChannel

	// Start another process thread to run in background
	// and signal when done
	go process(doneChannel)

	// Make channel non-blocking with select statement
	// This gives you the ability to continue executing
	// even if no message is waiting in the channel
	var readyToExit = false
	for !readyToExit {
		select {
		case done := <-doneChannel:
			log.Println("Done message received.", done)
			readyToExit = true
		default:
			log.Println("No done signal yet. Waiting.")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
