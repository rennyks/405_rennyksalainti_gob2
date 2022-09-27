package main

import (
	"fmt"
	"time"
)

func main() {

c1 := make(chan int, 3)

go func (c chan int)  {
	for i := 1; i <= 5; i++ {
	fmt.Println("func groutine starts sending data into the channel")
	c <- i
	fmt.Println("func groutine starts sending data into the channel")	
	}

	close(c)
}(c1)

fmt.Println("main groutine sleeps for 2 seconds")
time.Sleep(time.Second * 2)

for v := range c1 { // v = <- c1
	fmt.Println("main groutine received value from channel:", v)
}
}