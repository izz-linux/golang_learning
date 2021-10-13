package main

import (
	"fmt"
)

func main() {
	x := 44 % 40
	fmt.Println(x)

	y := 0
	for {
		y++
		// break exits the loop
		if y > 100 {
			break
		}
		// continue skips the rest of the loop and continues to the next iteration
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)

	}
}
