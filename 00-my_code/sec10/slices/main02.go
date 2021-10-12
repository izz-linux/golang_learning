package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// slicing a slice
	// includes start but up until and not including the end
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

}