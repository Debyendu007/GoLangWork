package main

import (
	"fmt"
)

func main() {
	arr := 42

	switch arr {
	case 1:
		fmt.Println("1. This is not 42")
	case 2:
		fmt.Println("2. This is not 42")
	case 3:
		fmt.Println("3. This is not 42")
	case 4:
		fmt.Println("4. This is not 42")
	case 5:
		fmt.Println("5. This is not 42")
	case (2 == 2):
		fmt.Println("6. This is not 42")
	case 7:
		fmt.Println("7. This is not 42")
	case 42:
		fmt.Println("8. This is 42")
	}
}
