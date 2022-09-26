package main

import (
	"fmt"
	"runtime"
)

func main() {
fmt.Println("main execultion started")

go firstPtocess(8)
secondProcess(8)

fmt.Println("no. of Goroutine:", runtime.NumGoroutine())
fmt.Println("main execution ended")
}

func firstPtocess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
	fmt.Println("i=", i)
	}
	fmt.Println("first process dunc ended")
}

func secondProcess(index int) {
	fmt.Println("second prpcess func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}