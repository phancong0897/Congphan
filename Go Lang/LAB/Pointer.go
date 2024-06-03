package main

import (
	"fmt"
)

// khai báo mảng
func modify(arr *[3]int) {
	(*arr)[1] = 97
	arr[2] = 99
}

func main() {
	// khai báo con trỏ
	a := 1997
	var b *int = &a
	fmt.Printf("Kiểu dữ liệu %T \n", b)
	fmt.Println("địa chỉ của b", b)
	// Giá trị zero của con trỏ
	var c *int
	if c == nil {
		fmt.Println("c is", c)
		c = &a
		fmt.Println("c after is", c)
	}
	// Tham chiếu ngược
	d := 1998
	e := &d
	fmt.Println("e address d:", e)
	fmt.Println("e value d:", *e)
	*e++
	fmt.Println("e vlaue + 1:", d)
	*e += 2
	fmt.Println("e value + 2:", d)
	// Sử dụng con trỏ trong slice
	f := [3]int{96, 98, 98}
	modify(&f)
	fmt.Println(f)

}
