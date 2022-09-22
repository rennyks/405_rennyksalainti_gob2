package main

import (
	"fmt"
	"strings"
)


func main()	{
	//cap function

	var fruits1 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits1 cap : ", cap(fruits1))//4
	fmt.Println("Fruits1 len : ", len(fruits1))//4

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]

	fmt.Println("Fruits1 cap : ", cap(fruits2))//4
	fmt.Println("Fruits1 len : ", len(fruits2))//3

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]
	fmt.Println("Fruits1 cap : ", cap(fruits3))//3
	fmt.Println("Fruits1 len : ", len(fruits3))//3

	//Creating a new backing array
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}
	newCars = append(newCars,cars[0:2]...)

	cars[0] = "nisan"
	fmt.Println("cars :", cars)
	fmt.Println("newCars:", newCars)
}