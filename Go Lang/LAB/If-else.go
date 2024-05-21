package main

import (
	"fmt"
)

func main() {
	a := 7
	if a%2 == 0 {
		fmt.Printf("%d la so chan", a)
	} else if a%3 == 0 {
		fmt.Printf("%d la so le", a)
	} else {
		fmt.Printf("%d khong chia het cho 2 va 3", a)
	}

}
