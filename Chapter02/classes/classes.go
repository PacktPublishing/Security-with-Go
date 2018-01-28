package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Doctor struct {
	Person         Person
	Specialization string
}

func main() {
	nanodano := Person{
		Name: "NanoDano",
		Age:  99,
	}

	drDano := Doctor{
		Person:         nanodano,
		Specialization: "Hacking",
	}

	fmt.Println(reflect.TypeOf(nanodano))
	fmt.Println(nanodano)

	fmt.Println(reflect.TypeOf(drDano))
	fmt.Println(drDano)
}
