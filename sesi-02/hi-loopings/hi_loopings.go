package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 3; i++{
	// 	fmt.Println("Angka", i)

	// }

	//var i = 1

	// for i <= 10{
	// 	fmt.Println("angka", i)
	// 	i++
	// }

	// for{
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++{
	// 	if i%2 == 1{
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }
		
	// for i := 0; i < 5; i++ {
	// for j := i; j < 5; j++ {
	// fmt.Print(j, " ")
	// 	}
	// fmt.Println()
	// }
	// outerLoop:
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Perulangan ke - ", i+1)
	// 	for j := 0; j < 3; j++ {
	// 		if i == 2 {
	// 			break outerLoop
	// 		}

	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}
	
	// var strings = [3] string {"arell", "nanda", "mailo"}
	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "manggo"

	// fmt.Printf("%#v\n", fruits)

	// var fruits = [3]string{"apel", "pisang", "mango"}
	// //01
	// for i, v := range fruits {
	// 	fmt.Printf("index : %d, value: %s\n", i, v)
	// }
	// fmt.Println(strings.Repeat("#", 25))

	// //02
	// for i := 0; i < len(fruits); i++{
	// 	fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	// }

	balances := [2][3]int{{5, 6, 7},{8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

