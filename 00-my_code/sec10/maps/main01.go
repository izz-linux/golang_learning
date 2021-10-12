package main

import "fmt"

func main() {
	m := map[string]int{
		"James":32,
		"Miss Moneypenny":27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	// the comma ok idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT22", v)
	}
}
