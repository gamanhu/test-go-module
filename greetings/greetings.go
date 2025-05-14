package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi and Hello %v. Welcome!", name)
	return message
}