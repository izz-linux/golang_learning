package main

import "fmt"

func main() {
	i := 1980
	for {
		if i <= 2021 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}
