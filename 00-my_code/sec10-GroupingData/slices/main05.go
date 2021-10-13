package main

import "fmt"

func main() {
	// using make (already know the size of a slice)
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	x = append(x, 3423)
	fmt.Println(cap(x))
	x = append(x, 3424)
	fmt.Println(cap(x))
	x = append(x, 3425)
	fmt.Println(cap(x))
	x = append(x, 3426)
	fmt.Println(cap(x))
}