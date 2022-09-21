package main

import "fmt"

func main() {
	var person map[string]string //deklarasi
	person = map[string]string{} //inisialisasi

	person["name"] = "Ariell"
	person["age"] = "23"
	person["address"] = "Manado"

	fmt.Println("Name :", person["name"])
	fmt.Println("Age :", person["age"])
	fmt.Println("Address :", person["address"])
}