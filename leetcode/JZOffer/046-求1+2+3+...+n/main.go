package main

import "fmt"

func SumRecursion(n int) int {
	ans := n
	if ans != 0 {
		ans += SumRecursion(n - 1)
	}
	return ans
}

func main() {
	fmt.Println(SumRecursion(3))
}
