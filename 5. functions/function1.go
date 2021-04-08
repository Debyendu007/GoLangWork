package main

import "fmt"

func main() {
	foo()
	bar("Raja")
	s1 := woo("Debyendu")
	x, y := mouse("Ian", "Fleming")

	fmt.Println(s1)
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s1 string) string {
	return fmt.Sprint("Hello from ", s1)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
