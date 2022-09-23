package main

import (
	"fmt"
	"strings"
)

func main()  {
	//mengembalikan sebuah nilai
	var names = []string{"renny", "jordan"}
	var printMsg = greet("heii", names)

	fmt.Println(printMsg)
}

func greet(msg string, names []string) string{
	var joinStr = strings.Join(names, " ")// function Join untuk menggabungkan nilai dengan tipe data string yg terdapat pada slice dan array

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}