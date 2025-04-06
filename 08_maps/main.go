package main

import "fmt"

func main() {
	//define map

	// emails := make(map[string]string)

	// //assign kv

	// emails["bob"] = "bob@gmail.com"

	// emails["sharon"] = "sharon@gmail.com"

	// emails["mike"] = "mike@gmail.com"

	// declare map and add key values
	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)

	fmt.Println(len(emails))

	fmt.Println(emails["bob"])

	//delete from map
	delete(emails, "bob")
	fmt.Println(emails)

}
