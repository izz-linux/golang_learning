package main

import "fmt"

func main() {

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

}
