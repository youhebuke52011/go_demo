package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	sb := []byte(strconv.Itoa(x))
	n := len(sb)
	l, r := 0, n-1
	for l < r {
		if sb[l] != sb[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
