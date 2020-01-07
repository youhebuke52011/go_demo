package main

import "fmt"

func reverseString(s []byte) {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l += 1
		r -= 1
	}
}

func main() {
	s := []byte("Hello")
	reverseString(s)
	fmt.Printf("%c", s)
}
