package main

import "fmt"

func main() {

	x := 10
	y := 10

	//if else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "red"

	// if color == "red" {
	// 	fmt.Println("The color is red")
	// } else if color == "blue" {
	// 	fmt.Println("The color is blue")
	// } else {
	// 	fmt.Println("The color is NOT red or blue")
	// }

//switch
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("The color is NOT red or blue")
			}
}
