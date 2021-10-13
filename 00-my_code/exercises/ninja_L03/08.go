package main

import "fmt"

func main() {
	switch {
	case (1 == 2):
		fmt.Println("This is actually false, but will print if 1 = 2")
	case (2 == 2):
		fmt.Println("2 will always equal true!")
	}
}