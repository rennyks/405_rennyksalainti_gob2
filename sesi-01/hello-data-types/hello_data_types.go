package main

import "fmt"

func main() {
	//number (integer) tipe data Integer
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("bilangan second : %T\n", second)
	fmt.Println("===============================")

	//belajar number(float)
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number : %f\n", decimalNumber)
	fmt.Printf("decimal Number : %.3f\n", decimalNumber)
	fmt.Println("===============================")

	//belajar Bool
	var condition bool = true
	fmt.Printf("is it premitted? %t \n", condition)
	fmt.Println("===============================")	
}