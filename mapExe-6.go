package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	x = []string{"Alabama", "Alaska", "Arizona", "Arakansas", "California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Lowa", "Kansas",
		"Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North California", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Califonia", "South Dakota", "Tannessee", "Texas",
		"Utsh", "Vermont", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
