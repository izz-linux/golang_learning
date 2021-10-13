package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// multidimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}