package main

import (
	"fmt"
)

func main() {

	var randomValues interface{}
	_ = randomValues

	randomValues = "jln sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"airel", "nanda"}

	fmt.Printf("output random value : %+v\n", randomValues)

	
	var v interface{}
	v = 20
	
	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}