package main

import "fmt"

func main(){

	// var person = map[string]string{
	// 	"name": "ariell",
	// 	"age":	"23",
	// 	"address":	"jalan sudirman",
	// }

	// fmt.Println("name :", person["name"])
	// fmt.Println("age :", person["age"])
	// fmt.Println("address :", person["address"])

	//looping with map
	var person = map[string]string{
		"name": "ariell",
		"age":	"23",
		"address":	"jalan sudirman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}