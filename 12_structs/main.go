package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
	
}

// greeting method (value reciever)
func (p Person) greet() string {

	return "hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// gets married (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else if p.gender == "F" {
		p.lastName = spouseLastName
	}

}

func main() {

	// //init person using struct
	person1 := Person{firstName: "Chika", lastName: "Ibeneme", city: "Kingston", gender: "M", age: 12}
	person2 := Person{firstName: "Jody", lastName: "Moulton", city: "Kingston", gender: "F", age: 13}

	// // alternative
	// // person2 := Person{"Chika", "Ibeneme", "Kingston", "M", 34}

	// fmt.Println(person1)
	// // fmt.Println(person2)

	// person1.age++
	// fmt.Println(person1)
	//person1.hasBirthday()
 person2.getMarried("Ibeneme")
 person1.getMarried("TestMarried")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
