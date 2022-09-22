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