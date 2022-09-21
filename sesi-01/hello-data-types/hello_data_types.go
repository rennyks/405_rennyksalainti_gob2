package main

import "fmt"

func main() {
	//default tipe data
	// var first = 89
	// var second = -12


	//inialisasi tipe data
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("bilangan second : %T\n", second)
	fmt.Println("===============================")

	//tipe data float
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number : %f\n", decimalNumber)
	fmt.Printf("decimal Number : %.3f\n", decimalNumber)
	fmt.Println("===============================")

	//tipe data boolean
	var condition bool = true
	fmt.Printf("is it premitted? %t \n", condition)
	fmt.Println("===============================")	

	//string
	var message string = "hallo renny"
	fmt.Print(message)
	fmt.Println("\n==========================")

	/* deklarasi string juga bisa dengan tanda grave accent/backticks (``) dengan 
	menggunakan backtics akan membuat semua karakter didalamnya tidak diescape*/
	var send_message = `Selamat datang di hactiv8. Salam kenal :).Mari belajar "scalable web service with go."`
	fmt.Printf(send_message)
	
}