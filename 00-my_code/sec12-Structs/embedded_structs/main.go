package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "here is a collision",
		ltk:   true,
	}

	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(sa1)
	fmt.Println(p1)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p1.first, p1.last, p1.age)

	// what about name collisions?  add "first" to secretAgent struct [wouldn't normally want to do this]
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
}
