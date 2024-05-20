package main

import (
	"fmt"
)

func main() {
	const (
		a = 50
		b = 25
		c = 16
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	const n = 25
	var name = n
	fmt.Printf("type %T value %v", name, name)
}
