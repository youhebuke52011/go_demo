package main

import "fmt"

type F func(int) int

func (f F) Fun() {
	fmt.Println("Fun1")
}

func Funn(i int) int {
	result := i + 1
	fmt.Println("Fun2")
	return result
}

func main() {
	t := F(Funn)

	t.Fun()
	fmt.Println(t(10))
}
