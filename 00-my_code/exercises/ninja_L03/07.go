package main

import "fmt"

func main() {
	myBD := 24
	if myBD == 20 {
		fmt.Println("My birthdate is the 20th!")
	} else if myBD == 21 {
		fmt.Println("My birthdate is the 21st!")
	} else if myBD == 22 {
		fmt.Println("My birthdate is the 22nd!")
	} else if myBD == 23 {
		fmt.Println("My birthdate is the 23rd!")
	} else if myBD == 24 {
		fmt.Println("My birthdate is the 24th!")
	} else {
		fmt.Println("I give...")
	}
}
