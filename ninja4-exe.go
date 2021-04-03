package main

import (
	"fmt"
)

func main() {
	year := 1991

	for {
		if year <= 2021 {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}
}
