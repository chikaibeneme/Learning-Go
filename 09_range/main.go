package main

import "fmt"

func main() {

	// ids := []int{33, 44, 62, 578, 84, 3, 4236, 34, 23}

	// //loop through ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// //not using index
	// for _, id := range ids {
	// 	fmt.Printf(" ID: %d\n", id)
	// }

	// //add ids together
	// sum := 0

	// for _, id := range ids {
	// 	sum += id
	// }
	// fmt.Printf("Sum: %d\n", sum)

	//range with  maps
	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
