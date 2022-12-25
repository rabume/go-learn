package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int32
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	// First way to populate a struct
	jim := person{
		firstName: "Jim",
		lastName:  "Craig",
		contactInfo: contactInfo{
			"jim.craig@gmail.com",
			4252,
		},
	}

	jim.updateName("jimmy")
	jim.print()

	// Second way to populate a struct
	var me person

	me.firstName = "Max"
	me.lastName = "Muster"
	me.contactInfo.email = "max.muster@gmail.com"
	me.contactInfo.zipCode = 4233

	me.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
