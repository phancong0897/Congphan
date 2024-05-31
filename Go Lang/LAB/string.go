package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for a, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, a)
	}
}

func main() {
	name := "Phan Văn Công"
	fmt.Println(name)
	printBytes(name)
	fmt.Println("\n")
	printChars(name)
	fmt.Println("\n")
	name1 := "Señor"
	printBytes(name1)
	fmt.Printf("\n")
	printChars(name1)
	printCharsAndBytes(name)
	nameslice := []byte{67, 111, 110, 103} // mã  dạng decimal
	name2 := string(nameslice)
	fmt.Println(name2)
	name3slice := []rune{0x43, 0x6f, 0x6e, 0x67}
	name3 := string(name3slice)
	fmt.Println(name3)
}
