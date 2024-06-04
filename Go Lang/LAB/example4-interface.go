// Triển khai nhiều interface
package main

import (
	"fmt"
)

type Summation interface {
	sum()
}

type Subtraction interface {
	sub()
}

type CalculateLeavesLeft interface {
	Cacul() int
}
type Employee struct {
	Name   string
	id     int
	wage   int
	reward int
	punish int
}

func (e Employee) sum() {
	fmt.Printf("Ten: %s, ID: %d, Luong tong: %d$ \n", e.Name, e.id, (e.wage + e.reward))
}

func (e Employee) sub() {
	fmt.Printf("Ten: %s, ID: %d, Luong tru: %d$ \n", e.Name, e.id, (e.wage - e.punish))
}

func (e Employee) Cacul() int {
	return e.wage + e.reward - e.punish
}

// Nhúng interface
type Information interface {
	Summation
	Subtraction
	CalculateLeavesLeft
}

func main() {
	e := Employee{
		Name:   "Congpv",
		id:     1,
		wage:   20,
		reward: 5,
		punish: 10,
	}
	var a Summation = e
	a.sum()
	var b Subtraction = e
	b.sub()
	var c CalculateLeavesLeft = e
	fmt.Printf("Tru luong : %d \n", c.Cacul())
	// Nhúng interface
	var infor Information = e
	infor.sum()
	infor.sub()
	fmt.Println("Tru:", infor.Cacul())
}
