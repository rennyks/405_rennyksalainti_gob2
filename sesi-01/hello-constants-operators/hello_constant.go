package main

import "fmt"

func main() {
	/*nilai constan atau konstanta adalah jenis variabel
	yang bersifat tetap, dan nilainya tdk dapat diubah*/
	const full_name string = "Renny Salainti"
	fmt.Printf("Hello %s", full_name)
	fmt.Println("\n=========================")

	// var value = 2 + 2*3
	// fmt.Println(value)	
	var value = ( 2 + 2 ) * 3
	fmt.Println(value)
	fmt.Println("\n==========================")

	
}