package main

import "fmt"

type Person struct {
	Name string
}

func NewPerson() Person {
	return Person{
		Name: "Anonymous",
	}
}

func main() {
	p := NewPerson()
	fmt.Println(p)
}
