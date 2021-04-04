package main

import (
	"fmt"
)

func main() {
	//

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("old: ", x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println("new: ", x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println("after append : ", x)
}
