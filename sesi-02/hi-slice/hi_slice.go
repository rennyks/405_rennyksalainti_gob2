package main

import "fmt"

func main() {
	//make function
	// var fruits = []string{"apple", "banana", "mango"}
	// _ = fruits

	// var fruits = make([]string, 3)
	// _= fruits
	// fmt.Printf("%#v", fruits)

	//sppend function
	// var fruits = make([]string, 3)
	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fmt.Printf("%#v", fruits)

	// fruits = append(fruits, "apple", "banana", "mango")
	// fmt.Printf("%#v", fruits)

	//append function with ellipsis
	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)
	// fmt.Printf("%#v", fruits1)

	//copy function
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elements => ", nn)

}