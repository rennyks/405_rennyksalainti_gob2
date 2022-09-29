package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	
	defer catchErr()

	
	var username string
	//var password len

	fmt.Println("Username:")
	fmt.Scanln(&username)
	
	fmt.Println("Password:")

	password, _ := gopass.GetPasswdMasked()

	if valid, err := validateLogin(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("error occured:", r)
	}else {
		fmt.Println("username & password correct!")
	}	
}


func validateLogin(username string,password []byte) (string, error) {
	pl := string(password)
	if username != "reni" && pl != "0987re" {
	return "", errors.New("is wrong!")
		
	}

	return "login!", nil
}