package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Printf("Dec = %d\tBin = %b\t%#X\n", a, a, a)

	b := a << 1
	fmt.Printf("Dec = %d\tBin = %b\t%#X\n", b, b, b)
}
