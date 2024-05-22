package main

import (
	"fmt"
)

func main() {
	var congpv [4]int
	congpv[0] = 1
	congpv[1] = 2
	fmt.Println("congpv:", congpv)
	fmt.Println("congpv2:", congpv[2])
	x := []int{1, 2, 3, 4, 5}
	fmt.Println("x:", x)
	a := [4]string{"e", "b", "c", "d"}
	fmt.Println("a :", a)
	b := a
	b[1] = "cong"
	fmt.Println("b:", b)
	fmt.Println("do dai b", len(b))

	for i, v := range x { //range
		fmt.Printf("%d bang gia tri: %d\n", i, v)
	}
	var z [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			z[i][j] = i + j
		}
	}
	fmt.Println("2d: ", z)
	var v [2][3]int
	v[0][0] = 1
	v[0][1] = 2
	v[0][2] = 3
	v[1][0] = 4
	v[1][1] = 5
	v[1][2] = 6
	fmt.Println("2d:", v)
}
