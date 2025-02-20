package main

import (
	"fmt"

	"github.com/tlylt/golib/hello"
)

func main() {
	msg := hello.Hello("world")
	fmt.Println(msg)
}
