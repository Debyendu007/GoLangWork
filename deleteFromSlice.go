package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("old: ", x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
