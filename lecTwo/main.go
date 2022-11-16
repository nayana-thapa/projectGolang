package main

import "fmt"

type ContactInfo struct {
	email   string
	pinCode int
}
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {

	var person Person
	person.firstName = "Bob"
	person.lastName = "Rana"
	person.contact = ContactInfo{
		email:   "bob@gmail.com",
		pinCode: 560050,
	}

	fmt.Printf("%+v", person)

	individual := Person{
		firstName: "Nayana",
		lastName:  "Thapa",
		contact: ContactInfo{
			email:   "nayana@gmail.com",
			pinCode: 560019,
		},
	}

	fmt.Printf("%+v", individual)

}
