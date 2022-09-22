package main

import (
	"fmt"
)

func main(){
	// looping over string (rune-by-rune)
	// str1 := "makan"
	// str2 := "mancai"

	// fmt.Printf("str1 byte lenght => %d\n", len(str1))
	// fmt.Printf("str2 byte lenght => %d\n", len(str2))

	//menggunakan function RuneCountlnString dari package utf8
	// fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	//menggunakan range loop
	str := "manca"

	for i, s := range str {
	fmt.Printf("index => %d, string => %s\n", i, string(s))
	}

	
}