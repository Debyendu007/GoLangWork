package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 11)

	for i := 0; i < len(x); i++ {
		x[i] = 100
	}

	fmt.Println(x)
}
