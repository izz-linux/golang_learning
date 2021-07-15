package main

import "fmt"

type izzType int
var x izzType

func main() {
	fmt.Println(x)
	fmt.Printf("%T",x)
	x = 42
	fmt.Println("\n",x)

}
