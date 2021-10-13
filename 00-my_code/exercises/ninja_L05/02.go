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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, val := range v.favIC {
			fmt.Println(val)
		}
	}

}
