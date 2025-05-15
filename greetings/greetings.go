package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi and Hello baba %v. Welcome!", name)
	return message
}