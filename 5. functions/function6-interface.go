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

/*
Any structure having speak methos is also of type human
A value can be of two type
*/
type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("PERSON >> I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("SECRET AGENT >> I was passed into bar", h.(secretAgent).first)
	}
}

type hotdog int

func main() {
	s1 := secretAgent{
		person: person{
			"Debyendu",
			"Das",
		},
		ltk: true,
	}

	s2 := secretAgent{
		person: person{
			first: "Palash",
			last:  "Das",
		},
		ltk: false,
	}

	p1 := person{
		first: "Mr. Raja",
		last:  "Das",
	}

	s1.speak()
	s2.speak()

	fmt.Println(p1)
	fmt.Printf("%T", s1)
	fmt.Println("\n")

	bar(s1)
	bar(s2)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
