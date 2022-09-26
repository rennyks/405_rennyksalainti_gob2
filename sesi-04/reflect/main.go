package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23 //
	var reflectValue = reflect.ValueOf(number)

	fmt.Printf("tipe value : %+v\n", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
	fmt.Printf("nilai variabel: %+v\n", reflectValue.Interface())
	
	}	
}