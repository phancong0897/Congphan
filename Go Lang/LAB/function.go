package main

import (
	"fmt"
)

func cacu(a, b, c, d int) (int, int, int, int) {
	var tong = a + b + c + d
	var tru = a + b - c + d
	var nhan = a * b
	var chia = b / c
	return tong, tru, nhan, chia
}
func main() {
	tong, tru, nhan, chia := cacu(8, 6, 3, 1)
	fmt.Printf("tong bang %d \n", tong)
	fmt.Printf("tru bang %d \n", tru)
	fmt.Printf("nhan bang %d \n", nhan)
	fmt.Printf("chia bang %d \n", chia)
}
