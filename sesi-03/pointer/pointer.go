package main

import (
	"fmt"
	"strings"
)

func main(){
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("seconNumber :", *secondNumber)
	fmt.Println("firstNumber (memori address) :", secondNumber)
	fmt.Println("==============================================")

	var firstPerson string = "renny"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("seconPerson :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")
	*secondPerson = "Done"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("seconPerson :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)
	fmt.Println("==============================================")

	var a int = 10

	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)
}
	func changeValue(number *int){
		*number = 20
	}
