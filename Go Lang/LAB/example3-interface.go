//Triển khai interface với con trỏ và giá trị
package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Person struct {
	age  int
	name string
}

func (p Person) Describe() {
	fmt.Printf("Name : %s, age: %d \n", p.name, p.age)
}

type Address struct {
	state string
	city  string
}

func (a *Address) Describe() {
	fmt.Printf("State: %s, city : %s \n", a.state, a.city)
}

func main() {
	var d2 Describer
	a := Address{"QL1A", "HaTinh"}
	d2 = &a
	d2.Describe()
	var p1 Describer
	b := Person{28, "Cong"}
	p1 = b
	p1.Describe()
	b1 := Person{27, "Dung"}
	p1 = &b1
	p1.Describe()

}
