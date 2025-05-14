package main

import  (
	"fmt"
	"example/greetings"
)

func main() {
	fmt.Println("Hello guys")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}	