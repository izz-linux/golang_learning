package main

import "fmt"

/*
type person struct {
	first string
	last string
	age int
}
 */

func main() {
	/* was like this:
	p1 := person {
		first: "James",
		last: "Bond",
		age: 32,
	}
	*/

	// this struct does NOT HAVE A NAME therefore 'anonymous struct'
	p1 := struct {
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}

	fmt.Println(p1)
}
