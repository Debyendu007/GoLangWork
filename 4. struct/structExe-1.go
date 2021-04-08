package main

import "fmt"

type person struct {
	first    string
	last     string
	flavours []string
}

func main() {
	p1 := person{
		first:    "Debyendu",
		last:     "Das",
		flavours: []string{"vanilla", "Chocolate", "Butter scotch"},
	}

	p2 := person{
		first:    "Raja",
		last:     "Bhattacharyya",
		flavours: []string{"nalengur", "aampanna", "black current"},
	}

	for _, v1 := range p1.flavours {
		fmt.Println(v1)
	}

	for _, v2 := range p2.flavours {
		fmt.Println(v2)
	}
}
