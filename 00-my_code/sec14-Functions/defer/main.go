package main

import "fmt"

func main() {
	// runs at the exit of contextual function
	defer foo()
	bar()
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}