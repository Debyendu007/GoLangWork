package main

import "fmt"

func main() {
	defer bar()
	defer foo()
	sum()
}

func foo() {
	fmt.Println("Welcome foo")
}

func bar() {
	fmt.Println("Welcome bar")
}

func sum() {
	fmt.Println("Welcome sum")
}
