package main

import "fmt"

func main() {

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
