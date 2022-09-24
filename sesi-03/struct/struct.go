package main

import (
	"fmt"
	"strings"
)
type Person struct {
	name string
	age int
}


type Employee struct{
	name string
	age int
	division string
	person Person
}

type Mployee struct{
	division string
	person Person
}

func main(){
	
	var employee Employee
	employee.name = "renny"
	employee.age = 21
	employee.division = "curiculum developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println("==========================================")


	var employee1 = Employee{}
	employee1.name = "renny"
	employee1.age = 23
	employee1.division = "developer"

	var employee2 = Employee{name: "ananda", age: 23, division: "finance"}

	fmt.Printf("Employee1 : %+v\n", employee1)
	fmt.Printf("Employee2 : %+v\n", employee2)
	fmt.Println("==========================================")

	var employee3 = Employee{name: "renny", age: 21, division: "developer"}
	var employee4 *Employee = &employee3

	fmt.Println("Employee3 name :", employee3.name)
	fmt.Println("Employee4 name :", employee4.name)

	fmt.Println(strings.Repeat("#", 20))
	employee4.name = "ananda"

	fmt.Println("Employee3 name :", employee3.name)
	fmt.Println("Employee4 name :", employee4.name)
	fmt.Println("==========================================")

	var employee5 = Mployee{}
	employee5.person.name = "renny"
	employee5.person.age = 21
	employee5.division = "developer"

	fmt.Printf("%+v\n", employee5)
	fmt.Println("==========================================")

	//anonymous struct tanpa isian field
	var employee6 = struct {
		person Person
		division string
	}{}
	employee6.person = Person{name: "renny", age: 21}
	employee6.division ="developer"

	//anonymous struct dengan pengisian field
	var employee7 = struct {
		person Person
		division string
	}{
		person: Person{name: "renny", age: 21},
		division: "Finance",
	}
	
	fmt.Printf("Employee6: %+v\n", employee6)
	fmt.Printf("Employee7: %+v\n", employee7)
	fmt.Println("==========================================")

	var people =[]Person{
		{name: "renny", age: 21},
		{name: "ryan", age: 22},
		{name: "mailo", age: 21},
	}
	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("===============================")

	var employee8 = []struct {
		person Person
		division string
	}{
		{person: Person{name: "Renny", age: 21}, division: "ui/uyx"},
		{person: Person{name: "Ananda", age: 20}, division: "curiculum developer"},
		{person: Person{name: "Mailo", age: 22}, division: "developer"},
	}
	for _, v := range employee8 {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("===============================")

}