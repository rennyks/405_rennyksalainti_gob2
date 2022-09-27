package main

import (
	"fmt"
	"time"
)

func main() {

c1 := make(chan int)

go func (c chan int)  {
	fmt.Println("func groutine starts sending data into the channel")
	c <- 10
	fmt.Println("func groutine starts sending data into the channel")	
}(c1)

fmt.Println("main groutine sleeps for 2 seconds")
time.Sleep(time.Second * 2)

fmt.Println("main groutine sleepss for 2 seconds")
d := <-c1
fmt.Println("main groutine data:", d)

close(c1)
time.Sleep(time.Second)
}