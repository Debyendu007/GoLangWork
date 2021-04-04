package main

import "fmt"

func main() {
	mp := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	if v, ok := mp["A"]; ok {
		fmt.Println(v)
		delete(mp, "A")
	}

	if v, ok := mp["E"]; ok {
		fmt.Println(v)
		delete(mp, "E")
	}

	if v, ok := mp["XXX"]; ok {
		fmt.Println(v)
		delete(mp, "XXX")
	}

	fmt.Println(mp)
}
