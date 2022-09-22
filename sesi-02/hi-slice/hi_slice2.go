package main

import (
	"fmt"
)

func main(){
	//copy function
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 =>"	, fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elements => ", nn)
}