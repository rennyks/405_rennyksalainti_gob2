package main

import (
	"fmt"
	"math"
)

func (c circle) volume() float64{
return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}

	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("circle value: %+v\n", value)
		fmt.Printf("circle volume %f", value.volume())
	}
}