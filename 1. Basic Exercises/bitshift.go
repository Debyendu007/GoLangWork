package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	// kb := 1024
	// kb = 1 << 10
	// mb = kb * 1024
	// gb = mb * 1024
)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b", gb, gb)
	// x := 2
	// fmt.Printf("%d\t\t%b\n", x, x)

	// y := x << 2
	// fmt.Printf("%d\t\t%b", y, y)

}
