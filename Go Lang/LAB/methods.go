package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
	city string
	address
}
type address struct { //Phương thức ẩn danh
	state string
}

func (a address) fulladress() {
	fmt.Printf("Full Address: %s \n", a.state)
}

// Cấu trúc methods
func (e Employee) Information() {
	fmt.Printf("Name: %s, age: %d , adress : %s \n", e.name, e.age, e.city)
}

// Con trỏ

type Employee1 struct {
	name string
	age  int
	stt  int
}

func (e Employee1) changename(newname string) {
	e.name = newname
}

func (e *Employee1) changeage(newage int) {
	e.age = newage
}

func (e *Employee1) changestt(newstt int) {
	e.stt = newstt
}

func main() {
	nv1 := Employee{
		name: "Cong",
		age:  28,
		city: "Ha Tinh",
		address: address{
			state: "QL1A",
		},
	}
	nv1.Information()
	nv1.fulladress() // Phương thức trường ẩn danh
	//Con trỏ
	e := Employee1{
		name: "Cong",
		age:  27,
		stt:  1,
	}
	fmt.Printf("Name nv2: %s \n", e.name)
	e.changename("CongPV")
	fmt.Printf("Namen change nv2 : %s \n", e.name)
	fmt.Printf("Age nv2: %d \n", e.age)
	(&e).changeage(28)
	fmt.Printf("Age change nv2: %d", e.age)
	e.changestt(2)
	fmt.Printf("\nAge change nv2: %d", e.stt)
}
