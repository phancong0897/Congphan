package main

import (
	"fmt"
)

type Employee struct {
	firtName, lastName string
	age                int
}

// Trường ẩn danh
type Employee1 struct {
	string
	int
}

// Cấu trúc lồng nhau
type Person struct {
	firtName, lastName string
	age                int
	address            Address
	sdt                //trường được khuyến khích
}
type Address struct {
	city  string
	state string
}
type sdt struct {
	sdt string
}

func main() {
	// khai báo cấu trúc
	nv1 := Employee{
		firtName: "Phan",
		age:      25,
		lastName: "Cong",
	}
	nv2 := Employee{"Nguyen", "thien", 1}
	fmt.Println("NV1:", nv1)
	fmt.Println("NV2", nv2)
	// Cấu trúc ẩn danh
	var nv3 Employee
	fmt.Println("NV3:", nv3)
	nv4 := Employee{
		firtName: "Trinh",
		lastName: "Quan",
	}
	fmt.Println("NV4:", nv4)

	// Truy cập riêng lẻ của cấu trúc
	fmt.Println("FirtName nv1:", nv1.firtName)
	fmt.Println("Age nv2", nv2.age)

	// Trỏ đến cấu trúc
	nv5 := &Employee{"Tu", "Anh", 23}
	fmt.Println("FirtName nv5", (*nv5).firtName)
	fmt.Println("LastName nv5", nv5.age)

	// Trường ẩn danh
	nv6 := Employee1{"Congpv", 27}
	fmt.Println("NV6:", nv6)

	var nv7 Person
	nv7.firtName = "Nguyen"
	nv7.lastName = "Ha"
	nv7.age = 50
	nv7.address = Address{
		city:  "HaTinh",
		state: "QL1A",
	}
	nv7.sdt = sdt{"0123456789"}
	fmt.Println("Name nv7:", nv7.firtName)
	fmt.Println("Age nv7:", nv7.age)
	fmt.Println("state nv7:", nv7.address.state)
	fmt.Println("sdt nv7: ", nv7.sdt) //trường được khuyến khích
}
