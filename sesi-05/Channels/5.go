package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		c1 <- "hello!"
	}()

	go func ()  {
		time.Sleep(1 * time.Second)
		c2 <- "salut!"
	}()

	for i := 1; i <= 2; i++ {
		select{
		case msg := <-c1:
			fmt.Println("received", msg)
		case msg1 := <-c2:
			fmt.Println("received", msg1)
		}
	}
}
