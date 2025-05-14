package main

import  (
	"fmt"
	"github.com/gamanhu/test-go-module/greetings"
)

func main() {
	fmt.Println("Hello guys")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}	