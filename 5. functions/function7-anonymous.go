package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("This is an anonymous function ", x)
	}(42)
}
