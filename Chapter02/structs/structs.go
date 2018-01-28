package main

import "fmt"

func main() {
	// Define a Person type. Both fields public
	type Person struct {
		Name string
		Age  int
	}

	// Create a Person object and store the pointer to it
	nanodano := &Person{Name: "NanoDano", Age: 99}
	fmt.Println(nanodano)

	// Structs can also be embedded within other structs.
	// This replaces inheritance by simply storing the
	// data type as another variable.
	type Hacker struct {
		Person           Person
		FavoriteLanguage string
	}
	fmt.Println(nanodano)

	hacker := &Hacker{
		Person:           *nanodano,
		FavoriteLanguage: "Go",
	}
	fmt.Println(hacker)
	fmt.Println(hacker.Person.Name)

	fmt.Println(hacker)
}
