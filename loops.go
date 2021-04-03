package main

import (
	"fmt"
)

func main() {
	a := 10

	for a > 1 {
		fmt.Println(a)
		a -= 1
	}

	b := 10

	for {
		if b == 1 {
			break
		}
		fmt.Println(b)
		b--
	}

	fmt.Println("Done.")
}
