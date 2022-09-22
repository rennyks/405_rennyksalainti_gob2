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
	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copied elements => ", nn)

	//slicing
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	
	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:]// sama dengan fruits1[:len(fruits1)]
	// fmt.Printf("%#v\n", fruits5)

	//Combining slicing and append
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	
	// fruits1 = append(fruits1[:3], "rambutan")
	// fmt.Printf("%#v", fruits1)

	//backing array
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}
	
	var  fruits2 = fruits1[2:4]
	fruits2[0] = "rambutan"

	fmt.Println("fruits =>", fruits1)
	fmt.Println("fruits2 => ", fruits2)

}