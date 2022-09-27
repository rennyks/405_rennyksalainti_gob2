package main

import "fmt"

func main() {

	c := make(chan string)

	students := []string{"airell", "mailo", "indah"}

	for _, v := range students {
		go func(student string) {
		fmt.Println("Student", student)
		result := fmt.Sprintf("hai, my name is %s", student)
		c <- result
		}(v)
	}
	for i := 1; i <=3; i++ {
	print(c)
}
	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("hi, my name is %s", student)
	c <- result
}

