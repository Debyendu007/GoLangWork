package main

import (
	"fmt"
)

type cType int

var x cType
var y int
var b bool

func main() {
	a := 42
	b := "42"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println(b)
	fmt.Println("\n")
	fmt.Println(a == (b))
}
