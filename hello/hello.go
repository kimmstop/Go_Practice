package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

//	message, err := greetings.Hello("") // error occurred
	 message, err := greetings.Hello("KMS") // error not ocurred
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}


