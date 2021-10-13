package main

import "fmt"

func main() {
	i := 42
	fmt.Printf("%d\t%b\t%#x\n", i, i, i)
	j := i << 1
	fmt.Printf("%d\t%b\t%#x\n", j, j, j)
}