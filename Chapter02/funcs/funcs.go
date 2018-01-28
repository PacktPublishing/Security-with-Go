package main

import "fmt"

// Function with no parameters
func sayHello() {
	fmt.Println("Hello.")
}

// Function with one parameter
func greet(name string) {
	fmt.Printf("Hello, %s.\n", name)
}

// Function with multiple params of same type
func greetCustom(name, greeting string) {
	fmt.Printf("%s, %s.\n", greeting, name)
}

// Variadic parameters, unlimited parameters
func addAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Function with multiple return values
// Multiple values encapsulated by parenthesis
func checkStatus() (int, error) {
	return 200, nil
}

// Define a type as a function so it can be used
// as a return type
type greeterFunc func(string)

// Generate and return a function
func generateGreetFunc(greeting string) greeterFunc {
	return func(name string) {
		fmt.Printf("%s, %s.\n", greeting, name)
	}
}

func main() {
	sayHello()
	greet("NanoDano")
	greetCustom("NanoDano", "Hi")
	fmt.Println(addAll(4, 5, 2, 3, 9))

	russianGreet := generateGreetFunc("Привет")
	russianGreet("NanoDano")

	statusCode, err := checkStatus()
	fmt.Println(statusCode, err)
}
