package main

import "fmt"

func main() {

	x := []int{2,3,5,7,11,13,17,19,23,29}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("\n%T",x)
}
