package main

import "fmt"

// add something to a map and range for looping over a map
func main() {
	m := map[string]int{
		"James":32,
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	// adding an element to a map
	m["todd"] = 33

	// ranging over a map
	for key, value := range m {
		fmt.Println(key, value)
	}


}
