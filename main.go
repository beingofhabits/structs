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
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Name",
		lastName:  "Name",
		contact: contactInfo{
			email:   "j@g.c",
			zipCode: 9444,
		},
	}

	fmt.Printf("%+v", jim)
}
