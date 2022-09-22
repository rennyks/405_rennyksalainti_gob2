package main

import "fmt"

func main(){

	var people =[]map[string]string{
		{"name": "ariel", "age": "23"},
		{"name": "nanda", "age": "23"},
		{"name": "mailo", "age": "23"},
	}

	for i, person := range people {
		fmt.Printf("index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}