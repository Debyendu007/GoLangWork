package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (42 <= 50)
	c := (42 >= 40)
	d := (42 != 42)
	e := (42 < 42)
	f := (42 > 42)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
