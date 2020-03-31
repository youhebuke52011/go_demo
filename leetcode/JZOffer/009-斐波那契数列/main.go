package main

import "fmt"

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	f0, f1 := 0, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = f0 + f1
		f0, f1 = f1, res
	}
	return res
}

func main() {
	fmt.Println(Fibonacci(5))
}
