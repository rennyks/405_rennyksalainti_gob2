package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("invoke with defer")
	fmt.Println("befor exiting")
	os.Exit(1)
}