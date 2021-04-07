package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Debyendu",
		last:  "Das",
		age:   29,
	}

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Raja",
		last:  "Das",
		age:   28,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
