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
		firstName: "Name",
		lastName:  "Name",
		contactInfo: contactInfo{
			email:   "j@g.c",
			zipCode: 9444,
		},
	}

	fmt.Printf("%+v", jim)
}
