package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooooo, James"}

	fmt.Println(x)
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println("Record: ", i)

		for j, v1 := range v {
			fmt.Printf("\tIndex position: %v\tvalue:\t%v\n", j, v1)
		}
	}
}
