package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("<%v> <%v> <%v>", x,y,z)
	fmt.Println(s)
	/*
	fmt.Println(x)
	fmt.Printf("%T", x)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println(z)
	fmt.Printf("%T", z)
	*/


}
