package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Debyendu":  32,
		"Raja":      88,
		"Palash":    22,
		"Debashish": 45,
	}

	fmt.Println(m)
	fmt.Println(m["Debyendu"])

	v, ok := m["Debyendu"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Debyendu"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}
