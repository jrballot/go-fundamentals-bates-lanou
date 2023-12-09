package main

import (
	"fmt"
)

func main() {
	emails := map[string]string{
		"foo":"foo@gmail.com",
		"bar":"bar@gmail.com",
		"matt":"matt@gmail.com",
	}

	fmt.Println("This is the list of emails we have: ")
	for key,value := range emails {
		fmt.Println(key," ",value)
	}

	// delete an item from the map
	delete(emails,"matt")
	fmt.Println("New email: ", emails)

	// check if the key exist
	email, ok := emails["matt"]
	if !ok {
		fmt.Println("Email doesn't exist for user \"matt\": ", email)
	} 
}