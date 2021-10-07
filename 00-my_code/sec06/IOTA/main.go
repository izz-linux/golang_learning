package main

import "fmt"

// iota starts at 0 and increments itself within a block
const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}