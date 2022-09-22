package main

import "fmt"

func main() {
	//looping over string (byte-by-byte)
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// fmt.Printf("index: %d, byte: %d\n", i, city[i])
	// }

	var city []byte = []byte{74, 97, 187, 97, 116, 116, 97}

	fmt.Println(string(city))
}