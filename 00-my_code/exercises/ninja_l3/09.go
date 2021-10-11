package main

import "fmt"

func main() {
	var favSport string = "Football"
	switch favSport {
	case "Futbol":
		fmt.Println("favorite sport is futbol")
	case "Baseball":
		fmt.Println("favorit sport is baseball")
	default:
		fmt.Printf("I couldn't guess you're favorite sport because you specified %s", favSport)
	}
}
