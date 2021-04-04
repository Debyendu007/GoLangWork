package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 0, 5)
	s2 := [5]int{31, 32, 33, 34}

	for i := 0; i < cap(s1); i++ {
		s1 = append(s1, i+35)
	}

	for i, v := range s1 {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n\n", s1)

	for i, v := range s2 {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", s2)
}
