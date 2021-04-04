package main

import (
	"fmt"
)

func main() {
	s2 := []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}

	for i, v := range s2 {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", s2)
}
