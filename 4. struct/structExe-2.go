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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k1, v1 := range m {
		fmt.Println("Value for\t", k1)

		for i1, v11 := range v1.flavours {
			fmt.Println("\t", i1, v11)
		}
	}

	for k2, v2 := range m {
		fmt.Println("Value for\t", k2)

		for i2, v21 := range v2.flavours {
			fmt.Println("\t", i2, v21)
		}
	}
}
