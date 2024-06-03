package main

import (
	"fmt"
)

type Tester interface {
	test()
}

type myfloat float64

func (m myfloat) test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface is Type %T, value %v\n", t, t)
}

// Interface rỗng

func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

// switch
func findtype(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("Type is string, value is : %s\n", i.(string))
	case int:
		fmt.Printf("Type is int, value is: %d\n", i.(int))
	case float64:
		fmt.Printf("Type is float64, values is: %f\n", i.(float64))
	default:
		fmt.Printf("Unknow type\n")
	}
}

func main() {
	var t Tester
	f := myfloat(1.65)
	t = f
	describe(t)
	t.test()
	// truyền giá trị vào interface rỗng
	s := "Congpv"
	describe1(s)
	b := 67.69
	describe1(b)
	str := struct {
		name string
	}{
		name: "Congpv",
	}
	describe1(str)
	//switch
	findtype("Congpv")
	findtype(55)
	findtype(98.98)
	findtype(0.1245678)
}
