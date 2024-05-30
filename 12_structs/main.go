package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseName string) {
	p.lastName = spouseName
}

func main() {
	//Init struct
	person := Person{
		firstName: "Emad",
		lastName:  "Siddiq",
		city:      "San Francisco",
		gender:    "m",
		age:       25,
	}

	//Alternative

	person_u := Person{"Someone", "Else", "San Fran", "f", 25}

	fmt.Println(person)
	fmt.Println(person_u)

	fmt.Println(person.firstName)
	person.age++
	fmt.Println(person.age)
	fmt.Println(person.greet())
	person.hasBirthday()
	fmt.Println(person.greet())
	fmt.Println(person_u.greet())
	person_u.getMarried("Siddiq")
	fmt.Println(person_u.greet())

}
