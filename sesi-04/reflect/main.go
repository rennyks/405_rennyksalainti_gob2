package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23 //
	var reflectValue = reflect.ValueOf(number)

	fmt.Printf("reflect value : %+V\n", number)
	fmt.Printf("tipe value : %+V\n", reflect.Value.Type())


	fmt.Printf("nilai variabel: %+V\n", reflectValue.Interface())

	
}