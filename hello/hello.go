package main

import (
	"fmt"
	"example.com/greetings"
)


func main(){
	message := greetings.Hello("test")
	fmt.Println(message)
}