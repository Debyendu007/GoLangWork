package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type Student struct {
	person Person
	roll   int
}

func main() {
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   22,
	}

	p2 := Person{
		first: "Debyendu",
		last:  "Das",
		age:   29,
	}

	st1 := Student{
		person: p1,
		roll:   18,
	}

	st2 := Student{
		person: p2,
		roll:   72,
	}

	fmt.Println("Person 1: ", p1)
	fmt.Println("Person 2: ", p2)

	fmt.Println("Student 1: ", st1)
	fmt.Println("Student 2: ", st2)
}
