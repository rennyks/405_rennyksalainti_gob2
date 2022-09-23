package main

import (
	"fmt"
)

func main()  {
	//varadic funvtion #2
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)

	fmt.Println("Result :", result)
}

func sum(numbers ...int) int{

	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
	
