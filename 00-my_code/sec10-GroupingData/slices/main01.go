package main

import "fmt"

func main() {
	// use of a composite literal
	// x := type{values}
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	/*
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	*/
	for index, value := range x{
		fmt.Println(index, value)
	}

}