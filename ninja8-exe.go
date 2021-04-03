package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This is not true")
	case true:
		fmt.Println("Now this is what I call a proper true value")
		fallthrough
	case true:
		fmt.Println("THESE is what I call a proper true value")
	}
}
