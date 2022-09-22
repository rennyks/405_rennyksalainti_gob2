package main

import "fmt"

func main () {
	//temporary variabel
	var currentYear = 2021
	
	if age := currentYear - 2001; age < 17 {
		
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else  {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	//switch
	// var score = 0
	// switch score {
	// case 8:
	// fmt.Println("Perfect")
	// case 7:
	// fmt.Println("awesome")
	// default:
	// fmt.Println("not bad")
	// }

	//switch with relational operators

	// switch{
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// default:{
	// 	fmt.Println("study harder")
	// 	fmt.Println("you need to learn more")
	// }
	// }

	//switch ( falthrough keyword)

	// switch{
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// fallthrough
	// case score < 5:
	// 	fmt.Println("it is ok, but please study harder")
	// default:{
	// 	fmt.Println("study harder")
	// 	fmt.Println("you don't have a good score yet")
	// }
	// }

	var score = 5
	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	}else{
		if score == 5{
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		}else{
			fmt.Println("you can do it")
			if score == 0{
				fmt.Println("TRY HARDER!")
			}
		}
	}

}