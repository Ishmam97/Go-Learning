package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// p1 := person{"Ishmam", "Solaiman"}
	// p2 := person{firstName: "Tanzila", lastName: "Maria"}
	var p3 person
	p3.firstName = "test"
	p3.lastName = "test2"
	p3.contact.email = "test@gmail.com"
	p3.contact.phone = "014214515"
	// fmt.Println(p1, p2)
	// fmt.Printf("%+v \n", p3)
	p4 := person{
		firstName: "hello",
		lastName:  "world",
		contact: contactInfo{
			email: "testmail",
			phone: "testphone",
		},
	}
	// p4Pointer := &p4
	// p4Pointer.updateName("ISSHMEMMEEME")
	p4.updateName("traaa")
	p4.toString()
}
func (p person) toString() {
	fmt.Printf("%+v \n", p)
}
func (p *person) updateName(firstname string) {
	(*p).firstName = firstname
}
