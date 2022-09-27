package main

import "fmt"

func main() {

	c := make(chan string)

go introduce("Airell", c)
go introduce("Nanda", c)
go introduce("Mailo", c)

msg := <-c
fmt.Println(msg)

msg2 := <-c
fmt.Println(msg2)

msg3 := <-c
fmt.Println(msg3)

close(c)

}

func introduce(student string, c chan string) {
result := fmt.Sprintf("hai, my name is %s", student)

c <- result
}