package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck {
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println("--------------")

	fmt.Printf("True/False, the truck is 4 wheel drive: %v\n", t.fourWheel)
	fmt.Printf("The color of our sedan is %s\n", s.color)

}
