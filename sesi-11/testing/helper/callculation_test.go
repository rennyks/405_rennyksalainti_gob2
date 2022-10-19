package helper

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("result should be 20")
	}
}