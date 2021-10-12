package main

import "fmt"

// delete something from a map
func main() {
	m := map[string]int{
		"James":32,
		"Miss Moneypenny":27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	// can delete something that doesn't exist:
	delete(m, "Ian Fleming")
	fmt.Println(m)

	// be sure you actually delete something with comma ok idiom
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	// now it is empty
	fmt.Println(m)

}