package main

import "fmt"

func main()  {
	//01
	greet("renny", 21)
	// greet("renny", "manado")
}

func greet(name string, age int8)  {
	fmt.Printf(" hello there! my name is %s and i'm %d years old", name, age	)
	
}
	//02

func greet(name, address string)  {
	fmt.Println("hello there! my name is", name)
	fmt.Println("i live in", address)
}