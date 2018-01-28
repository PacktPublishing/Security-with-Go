package main

import "fmt"

// Define a custom type that will
// be used to satisfy the error interface
type customError struct {
	Message string
}

// Satisfy the error interface
// by implementing the Error() function
// which returns a string
func (e *customError) Error() string {
	return e.Message
}

// Sample function to demonstrate
// how to use the custom error
func testFunction() error {
	if true != false { // Mimic an error condition
		return &customError{"Something went wrong."}
	}
	return nil
}

func main() {
	err := testFunction()
	if err != nil {
		fmt.Println(err)
	}
}
