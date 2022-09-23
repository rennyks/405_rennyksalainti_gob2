package main

import (
	"fmt"
	"math"
)

func main()  {
	var diameter float64 = 15
	var  area, circumferense = calculate(diameter)

	fmt.Println("area :", area)
	fmt.Println("Circumferense :", circumferense)
}

//predefined return value
func calculate(d float64) (area float64, circumferense float64) {
	//menghitung luas
	area = math.Pi * math.Pow(d/2, 2)

	//menghitung keliling
	circumferense = math.Pi * d
	
	return
	
}