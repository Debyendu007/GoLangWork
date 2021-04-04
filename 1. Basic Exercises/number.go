package main

import (
	"fmt"
)

func main() {
	// s := "Hello World"
	// fmt.Println(runtime.GOOS)
	// fmt.Println(runtime.GOARCH)

	// fmt.Printf("%T\n", s)
	// fmt.Println("%s\n", s)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%#U\n ", s[i])
	// }

	// for i, v := range s {
	// 	fmt.Printf("at index position %d we have hex %#x\n", i, v)
	// }

	cnt := "H"
	bt := []byte(cnt)

	fmt.Println("cnt: ", cnt)
	fmt.Println("bt: ", bt)

	num := bt[0]
	fmt.Printf("num: decimal %d\n", num)
	fmt.Printf("num: binary %b\n", num)
	fmt.Printf("num: hex %#X", num)
}
