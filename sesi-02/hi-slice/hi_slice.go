package main

import "fmt"

func main() {
	//make function
	// var fruits = []string{"apple", "banana", "mango"}
	// _ = fruits

	var fruits = make([]string, 3)
	_= fruits
	fmt.Printf("%#v", fruits)
	
	//backing array
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	
	var  fruits2 = fruits1[2:4]
	fruits2[0] = "rambutan"

	fmt.Println("fruits =>", fruits1)
	fmt.Println("fruits2 => ", fruits2)

}