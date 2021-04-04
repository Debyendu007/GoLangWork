package main

import "fmt"

func main() {
	s2 := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := s2[2:7]
	b := s2[7:]
	c := s2[4:9]
	d := s2[3:8]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
