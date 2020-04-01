package main

import (
	"fmt"
)

/**
题目：“student. a am I”  -->  “I am a student.”
思路：找规律，先把每个单词翻转一次，然后整条句子再翻转一次
*/

func ObliqueWordSequence(s []byte) []byte {
	l := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			swap(s, l, i-1)
			l = i + 1
		}
	}
	swap(s, 0, len(s)-1)
	return s
}

func swap(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l += 1
		r -= 1
	}
}

func main() {
	res := ObliqueWordSequence([]byte("student. a am I"))
	fmt.Printf("%s", res)
}
