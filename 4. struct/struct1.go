package main

import "fmt"

type Person struct {
	first string
	last  string
}

func main() {
	p1 := Person{
		first: "James",
		last:  "Bond",
	}

	p2 := Person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
