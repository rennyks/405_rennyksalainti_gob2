package main

import (
	"fmt"
)

func main() {

	//first way of looping
	for i := 0; i < 3; i++{
		fmt.Println("Angka", i)

	}
	//second way of looping
	var i = 1
	
	for i <= 10{
		fmt.Println("angka", i)
		i++
	}
	
	//third way of looping
	for{
		fmt.Println("Angka", i)
		i++
		if i == 3 {
			break
		}
	}

	
	
	
}

