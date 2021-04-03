package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Dibyendu"}
	fmt.Println(jb)

	mp := []string{"Raja", "Das", "Palash", "Debashish"}
	fmt.Println(mp)

	mds := [][]string{jb, mp}
	fmt.Println(mds)

	mds1 := [][]string{mp, jb}
	fmt.Println(mds1)

	mds2 := [][][]string{mds, mds1}
	fmt.Println(mds2)

}
