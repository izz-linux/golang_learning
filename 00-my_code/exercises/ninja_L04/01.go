package main

import "fmt"

func main() {

	x := [5]int{2,3,5,7,11}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("\n\tType of Array is of %T", x)

}
