package main

import(
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("KMS")
	fmt.Println(message)
}


