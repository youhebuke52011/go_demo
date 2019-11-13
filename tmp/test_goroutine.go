package main

import (
	"fmt"
)

var ch = make(chan bool)

func hello() {
	fmt.Println("hello world")
	close(ch)
}

func main() {
	go hello()
	<- ch
}
