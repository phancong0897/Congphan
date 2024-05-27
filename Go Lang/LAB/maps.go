package main

import (
	"fmt"
)

func main() {
	name := map[string]int{
		"congpv1": 27052024,
		"hau1":    27052024,
	}
	name["congpv"] = 120897
	name["Hau"] = 100698
	name["Minh"] = 250823
	fmt.Println("ten va so :", name)
	ten := "congpv"
	fmt.Println("Ten la", ten, "so", name[ten])
}
