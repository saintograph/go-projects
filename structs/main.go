package main

import (
	"fmt"
)

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Fellow",
		contact: contact{
			email:   "something@gmail.com",
			zipCode: 99000,
		},
	}
	jim.updateName("Mary")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
