package main

import "fmt"

func main() {
	/* operator (arithmetic operators)*/
	// var value = 2 + 2*3
	// fmt.Println(value)	
	var value = ( 2 + 2 ) * 3
	fmt.Println(value)
	fmt.Println("\n==========================")

	/*Operators (Relational Operators)*/
	var firstCondition bool = 2 < 3
	var secondCondition bool = "renny" == "Renny"
	var thirdCondition bool = 10 != 2.3
	var fourthcondition bool = 11 <= 11

	fmt.Println("firstCondition:", firstCondition)
	fmt.Println("secondCondition:", secondCondition)
	fmt.Println("thirdCondition:", thirdCondition)
	fmt.Println("fourthCondition:", fourthcondition)
	fmt.Println("\n=====================================")

	/*Operators (Logical Operators*/
	var right = true
	var left = false
	
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
	fmt.Println("\n======================================")

}