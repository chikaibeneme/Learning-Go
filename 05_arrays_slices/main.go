package main

import "fmt"

func main() {

	//arrays

	//var fruitArr [2]string

	//assign values

	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	// //declare and assign values
	// fruitArr := [2]string{"apple", "orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])
	fruitSlice := []string{"apple", "orange", "grape", "banana"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
