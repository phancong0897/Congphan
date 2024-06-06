/*
Nhập vào a,b,c . Kiểm tra xem a,b,c có tạo nên được tam giác không?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your a:")
	astr, _ := reader.ReadString('\n')
	fmt.Print("Enter your b:")
	bstr, _ := reader.ReadString('\n')
	fmt.Print("Enter your c:")
	cstr, _ := reader.ReadString('\n')

	// Loại bỏ các khoảng trắng và ký tự xuống dòng
	astr = strings.TrimSpace(astr)
	bstr = strings.TrimSpace(bstr)
	cstr = strings.TrimSpace(cstr)
	// Chuyển đổi chuỗi sang float 64
	a, err1 := strconv.ParseFloat(astr, 64)
	b, err2 := strconv.ParseFloat(bstr, 64)
	c, err3 := strconv.ParseFloat(cstr, 64)
	// Kiểm tra chuyển đổi
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Lỗi đầu vào không hợp lệ, vui lòng nhập lại")
	}
	if a+b > c && b+c > a && c+a > b {
		fmt.Println("a,b,c là 3 cạnh của hình chữ nhật")
	} else {
		fmt.Println("a,b,c không phải 3 canh của hình chữ nhật")
	}
	fmt.Println()(check(3, 4, 5))
}
