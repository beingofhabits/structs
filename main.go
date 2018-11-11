package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Malkovich",
		contactInfo: contactInfo{
			email:   "j@g.c",
			zipCode: 9444,
		},
	}

	jim.updateName("Jimmy")

	jim.print()
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
