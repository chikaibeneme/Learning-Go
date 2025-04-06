package main

import "fmt"

func main() {

	//Using var

	var age int32 = 29
	const isCool = true
	//isCool = false

	//shorthand
	name := "Chika"
	size := 1.3

	fName, email := "Chika", "chikaibeneme6@gmail.com"

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", isCool)
	fmt.Println(fName, email)

}
