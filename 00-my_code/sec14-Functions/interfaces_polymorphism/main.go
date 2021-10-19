package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// interfaces allow us to define behavior and allow us to do polymorphism
// any other type which has the method speak() is also of type human
// a value can be of more than one type!!!!
type human interface {
	// hey baby, if you have this method, then you're my type!
	speak()
}

// func (r receiver) identifier(parameters) (return(s)) {...}
/* when you have a receiver it's going to attach this function to the type
   so any value of that type then has access to this method
*/
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func bar(h human) {
	/* we can't do the following:
	 fmt.Println("I was passed into bar", h.first)
	 because we don't know if h is secretAgent of person, so we have a special switch on type:
	 */
	switch h.(type) { // assertion, assert on type
	case person:
		fmt.Println(h.(person).first, h.(person).last, "is a person")
	case secretAgent:
		fmt.Println(h.(secretAgent).first, h.(secretAgent).last, "is a secretAgent")
	default:
		fmt.Println(h, "is an unknown type")
	}
	fmt.Println("I was passed into bar\n\n", h)
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

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}

	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}



