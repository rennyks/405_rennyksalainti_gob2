package main

import "fmt"

func main()  {
	greet("airel", "sudirman")
}

func greet(name, address string)  {
	fmt.Println("hello there! my name is", name)
	fmt.Println("i live in", address)
}