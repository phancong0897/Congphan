package main

import (
	"fmt"
)

func main() {
	a := 9
	switch a {
	case 1, 3, 5, 7, 9:
		fmt.Println(" so le")
	case 2, 4, 6, 8:
		fmt.Println("so chan")
	default:
		fmt.Println("sai so")
	}
	b := 12
	switch {
	case b >= 0 && b <= 10:
		fmt.Println("so be hon 10")
	case b > 10:
		fmt.Println("so lon hon 10")
	}
	c := 9
	switch {
	case c%2 == 0:
		fmt.Printf("chia het cho 2 \n")
		fallthrough
	case c%3 == 0:
		fmt.Printf("chia het cho 3 \n")
	}
}
