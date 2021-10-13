package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// append to a slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	// ... is a variadic parameter
	x = append(x, y...)

	fmt.Println(x)

	// delete from a slice (will take out at index 2 and 3 or the 7 and 8 values
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}