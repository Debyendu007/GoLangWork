package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i - 65 + 1)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
