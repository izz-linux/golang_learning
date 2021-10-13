package main

import "fmt"

func main() {
	/*
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		// no default fallthrough in Go
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		// next after a fallthrough will print (true or not)
		fmt.Println("not true 1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
	default:
		fmt.Println("this is default")
	}
	 */

	n := "Bond"
	switch n {
	case "Moneypenny", "bond", "Do No":
		fmt.Println("multiple cases")
	case "Bond":
		fmt.Println("bond james")
	}

}
