package main

import (
	"fmt"
)

// Hàm bất định để tính tổng các số nguyên
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Gọi hàm với các tham số khác nhau
	fmt.Println(sum(1, 2, 3))        // Output: 6
	fmt.Println(sum(10, 20, 30, 40)) // Output: 100
	fmt.Println(sum())               // Output: 0
}
