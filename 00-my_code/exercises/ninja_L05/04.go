package main

import "fmt"

func main() {
	// this struct does NOT HAVE A NAME therefore 'anonymous struct'
	house := struct {
		houseNum int
		street string
		city string
		state string
		zip int
	}{
		houseNum: 6908,
		street: "Ridge Crossings Pkwy",
		city: "Hoover",
		state: "AL",
		zip: 35244,
	}

	fmt.Printf("My old address was an apartment at:\n")
	fmt.Printf("%d %s\n%s, %s %d", house.houseNum, house.street, house.city, house.state, house.zip)
}
