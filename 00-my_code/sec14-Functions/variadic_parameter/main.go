package main

import "fmt"

func main() {
	x := foo(2,3,4,5,6,7,8,9)
	fmt.Println(x)
}

// func (r receiver) identifier(parameters) (return(s)) {...}

// VARIADIC PARAMETER MUST BE THE FINAL PARAMETER
func foo(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}