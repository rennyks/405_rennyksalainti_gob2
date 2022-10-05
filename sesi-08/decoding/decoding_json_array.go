package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func main() {
	var jsonString = ` [
		{
		"full_name": "airel jordan",
		"email": "airel@gmail.com",
		"age": 23
		},

		{
		"full_name": "ananda",
		"email": "ananda@gmail.com",
		"age": 23
		}
		]
		`
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i , v := range result {
		fmt.Printf("index %d: %+v\n", i+1, v)
	}
}