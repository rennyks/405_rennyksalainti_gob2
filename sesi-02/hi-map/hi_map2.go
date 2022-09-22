package main

import "fmt"

func main(){
	
	var person = map[string]string{
		"name": "ariell",
		"age":	"23",
		"address":	"jalan sudirman",
	}
	
	// //deleting value
	// fmt.Println("Before deleting : ", person)
	// delete(person, "age")
	// fmt.Println("After deleting : ", person)

	//detecting a value
	value, exist := person["age"]

	if exist{
		fmt.Println(value)
	}else{
		fmt.Println("value does'nt exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist{
		fmt.Println(value)
	}else{
		fmt.Println("value has been dletend")
	}
}