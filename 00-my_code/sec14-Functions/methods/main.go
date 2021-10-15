package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}
	sa1.speak()
	sa2.speak()
}

// func (r receiver) identifier(parameters) (return(s)) {...}
/* when you have a receiver it's going to attach this function to the type
   so any value of that type then has access to this method
 */
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}