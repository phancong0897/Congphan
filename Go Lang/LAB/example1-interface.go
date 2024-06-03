package main

import (
	"fmt"
)

type salarycacul interface {
	Caculate() int
}
type permanent struct {
	wage   int
	reward int
}

type contract struct {
	reward int
}

func (p permanent) Caculate() int {
	return p.wage + p.reward
}

func (c contract) Caculate() int {
	return c.reward
}

func totalExpense(s []salarycacul) {
	expense := 0
	for _, v := range s {
		expense = expense + v.Caculate()
	}
	fmt.Printf("Total Expense %d$", expense)
}

func main() {
	cong := permanent{10000, 5000}
	hau := permanent{6000, 2000}
	dung := contract{5000}
	total := []salarycacul{cong, hau, dung}
	totalExpense(total)
}
