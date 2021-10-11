package main

import "fmt"

func main() {

	/*
	if true {
		fmt.Println("001")
	}
	if !false {
		fmt.Println("002")
	}
	if false {
		fmt.Println("don't print this")
	}

	if x := 42; x == 42 {
		fmt.Println("hell yea!")
	}
	*/

	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}

}
