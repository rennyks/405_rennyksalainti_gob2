package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {

	//username dengan tipe data string
	var username string

	//untuk input username
	fmt.Println("Username:")
	fmt.Scanln(&username)
	
	//untuk input password
	fmt.Println("Password:")
	paswd, _ := gopass.GetPasswdMasked()


	
	//melakukan validasi pada username dan password
	if valid, err := validateLogin(username, paswd); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}


//fungction validateLogin mengatur username & paswd
func validateLogin(username string, password []byte) (string, error) {
	ps := string(password[:])
	if username != "reni" && ps != "0987re" {
	return "", errors.New("is wrong!")
		
	}
	return "login!", nil
}