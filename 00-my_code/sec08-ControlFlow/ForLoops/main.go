package main

import "fmt"

func main() {
	// for init; condition; post {}

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}

	// according to language spec, can also do this (for is controlled by at least a boolean eval):
	x := 7
	for x < 13 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
	// ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
	x = 1

	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x ++
	}

}