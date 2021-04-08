package main

import "fmt"

func main() {
	str1 := struct {
		firstName string
		friends   map[string]int
		favDrink  []string
	}{
		firstName: "Debyendu Das",
		friends: map[string]int{
			"Debyendu":  555,
			"Raja":      445,
			"Palash":    778,
			"Debashish": 569,
		},
		favDrink: []string{"Coke", "ThumsUp", "7UP", "Sprite"},
	}

	fmt.Println(str1)
}
