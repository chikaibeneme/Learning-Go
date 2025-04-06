package main

import "fmt"

func main() {

	// //long method for a for loop

	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }

	// //short method

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	//Fizzbuzz
	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Printf("fizzbuzz\n")
		} else if i%3 == 0 {
			fmt.Printf("fizz\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz\n")
		} else {

			fmt.Println(i)

		}
	}

}
