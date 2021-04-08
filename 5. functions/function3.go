//Variadic parameters

package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum1(xi...)
	fmt.Printf("\tX >> The sum is %v", x)

	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("\tSUM >> The sum is %v", sum)

	xa := sum2(xi)
	fmt.Printf("\tSUM2 >> The sum is %v", xa)
}

func sum(x ...int) int {
	// fmt.Println(x) >> []
	// fmt.Printf("%T\n", x) >> []int

	sum := 0

	fmt.Println("For function SUM")

	for i, v := range x {
		sum += v
		fmt.Printf("\tFor item in index position %v we are now adding %v to the total which is now %v\n", i, v, sum)
	}

	return sum
}

func sum1(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	fmt.Println("For function SUM1")

	for i, v := range x {
		sum += v
		fmt.Printf("\tFor item in index position %v we are now adding %v to the total which is now %v\n", i, v, sum)
	}

	return sum
}

func sum2(x []int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	fmt.Println("For function SUM2")

	for i, v := range x {
		sum += v
		fmt.Printf("\tFor item in index position %v we are now adding %v to the total which is now %v\n", i, v, sum)
	}

	return sum
}
