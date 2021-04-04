package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Icce cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("Key: ", k)

		for i, val := range v {
			fmt.Printf("\tIndex:\t%v\t\tValue:\t%v\n", i, val)
		}
	}
}
