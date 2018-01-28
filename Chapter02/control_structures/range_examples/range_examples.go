package main

import "fmt"

func main() {
	intSlice := []int{2, 4, 6, 8}
	for key, value := range intSlice {
		fmt.Println(key, value)
	}

	myMap := map[string]string{
		"d": "Donut",
		"o": "Operator",
	}

	// Iterate over a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Iterate but only utilize keys
	for key := range myMap {
		fmt.Println(key)
	}

	// Use underscore to ignore keys
	for _, value := range myMap {
		fmt.Println(value)
	}
}
