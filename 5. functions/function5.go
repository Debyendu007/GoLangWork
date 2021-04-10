package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last)
}

func main() {
	s1 := secretAgent{
		person: person{
			"Debyendu",
			"Das",
		},
		ltk: true,
	}

	s1.speak()
}
