package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	res := make([]int, n+1)
	t := 1
	for i := n - 1; i >= 0; i-- {
		res[n] = (digits[i] + t) % 10
		n--
		t = (digits[i] + t) / 10
	}
	if t != 0 {
		res[n] = t
	}
	if res[0] == 0 {
		return res[1:]
	}
	return res
}

func main() {
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}
