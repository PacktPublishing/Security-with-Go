package main

import "fmt"

func main() {
	// Create a nil slice
	var mySlice []byte

	// Create a byte slice of length 8 and max capacity 128
	mySlice = make([]byte, 8, 128)

	// Maximum capacity of the slice
	fmt.Println("Capacity:", cap(mySlice))

	// Current length of slice
	fmt.Println("Length:", len(mySlice))
}
