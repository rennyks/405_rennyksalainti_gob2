package main

import "fmt"

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings) 

	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "manggo"

	// fmt.Printf("%#v\n", fruits)

	var fruits = [3]string{"apel", "pisang", "mango"}
	
	// //01
	// for i, v := range fruits {
	// 	fmt.Printf("index : %d, value: %s\n", i, v)
	// }
	// fmt.Println(strings.Repeat("#", 25))

	// //02
	for i := 0; i < len(fruits); i++{
		fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	}
}
