package main

import "fmt"

func main() {
	
	//belajar variabel with type data
	var namaLengkap string
	var umur int = 99

	namaLengkap = "Renny"
	umur = 21

	fmt.Println("Hello ", namaLengkap)
	fmt.Println("Umur ", umur)
	fmt.Println("=======================")

	//belajar variabel without type data
	var namaAnda = "Anonymous"
	var umurAnda = true

	fmt.Printf("%T, %T\n", namaAnda, umurAnda)
	println("=================================")

	//belajar short declaration
	namaSaya := "Renny"
	umurSaya := 21

	fmt.Printf("%T, %T\n", namaSaya, umurSaya)
	fmt.Println("===============================")

	//belajar multiple variabel declaration
	var one, two, three string = "1", "2", "3"
	var angkaOne, angkaTwo, angkaThree int = 1, 2, 3
	fmt.Println("test data string multiple > ", one, two, three)
	fmt.Println("test data int multiple > ", angkaOne, angkaTwo, angkaThree)
	fmt.Println("============================")

	//belajar multiple variabel declarasion short
	var oneData, twoData, threeData = 1, true, "Renny"
	dataOne, dataTwo, dataThree := true, 55.5, "Renny"
	fmt.Println("test data string multiple > ", oneData, twoData, threeData)
	fmt.Println("test data int multiple > ", dataOne, dataTwo, dataThree)
	fmt.Println("============================")

	//belajar underscore variabel
	var testVariabel string
	var oneName, twoName, alamatSaya, dataUmur = "Renny", "Salainti", "Manado", 21
	_, _, _, _, _ = testVariabel, oneName, twoName, alamatSaya, dataUmur
	fmt.Printf("test underscore variabel > %T", testVariabel)
	fmt.Printf("halo nama saya %s, umur saya %d, saya beralamat di %s\n", oneName, dataUmur, alamatSaya)
}