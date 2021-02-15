package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"ABC", "DEF", "KMD"}

//	message, err := greetings.Hello("") // error occurred
//	message, err := greetings.Hello("KMS") // error not ocurred
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}


