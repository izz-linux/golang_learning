package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIC []string
}

func main() {

	p1 := person {
		firstName: "Izz",
		lastName: "Noland",
		favIC: []string{"Cookie Dough", "Reeses"},
	}

	p2 := person {
		firstName: "Anna",
		lastName: "Whitney",
		favIC: []string{"Vanilla", "Coffee"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, "loves the following ice creams:")
	fmt.Printf("\t")
	for _,v := range p1.favIC {
		fmt.Printf("%s,", v)
	}

	fmt.Println("\n", p2.firstName, p2.lastName, "loves the following ice creams:")
	fmt.Printf("\t")
	for _,v := range p2.favIC {
		fmt.Printf("%s, ", v)
	}
}
