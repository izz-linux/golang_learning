package main

import "fmt"

func main() {
	foo()
	bar("Izz")
	s1 := woo(`Moneypenny`)
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(str string) string {
	return fmt.Sprintf("Hello from woo, %s", str)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello"`)
	b := false
	return a, b
}