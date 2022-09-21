package main

import "fmt"

func main() {
	var currentYear = 2021
	
	if age := currentYear - 2001; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	// var score = 10
	// switch score {
	// case 8:
	// fmt.Println("Perfect")
	// case 7:
	// fmt.Println("awesome")
	// default:
	// fmt.Println("not bad")
	// }

	// var score = 6

	// switch{
	// 	case score
	//}
}

