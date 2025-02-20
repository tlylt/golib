package hello

import "fmt"

func Hello(input string) string {
	msg := fmt.Sprintf("Hello there, %s!", input)
	return msg
}
