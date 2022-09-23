package main

import (
	"fmt"
	"strings"
)

func main() {

	var studentLists = []string{"airel", "nanda",}
	var find = findStudent{studentLists}
	fmt.Println(find("airell"))
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student
			}
		}
	}
}