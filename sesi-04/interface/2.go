package main

import "fmt"

func main() {

	//empty interface

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

	rs := []interface{}{1, "ariell", true, 2, "ananda", true}
	rm := map[string]interface{}{
		"nama":   "airell",
		"status": true,
		"age":    23,
	}
	_, _ = rs, rm
}