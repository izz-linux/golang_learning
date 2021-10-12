package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xp := [][]string{jb, mp}

	fmt.Println(xp)

	for i, v := range xp {
		fmt.Println("record: ", i)
		for j, value := range v {
			fmt.Printf("\tindex position%v\t value: %v\n", j, value)
		}
	}

}
