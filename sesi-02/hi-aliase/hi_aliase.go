package main

import "fmt"

func main() {
	//deklarasi variabel dengan tipe data units
	var a uint8
	var b byte //byte adalah alias dari tipe data uint8

	b = a //no error, karena byte memiliki tipe data yang sama dengan uint8
	_ = b

	/*mendeklarasikan tipe data alias bernama second
	type nama_alias = nama_tipe_data*/
	type second = uint
	var hour second = 3600

	fmt.Printf("hour type : %T\n", hour) // => hour type: uint
}