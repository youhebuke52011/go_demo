package main

import "fmt"

// 反转该字符串中的元音字母

func reverseVowels(s string) string {
	sb := []byte(s)
	n := len(s)
	l, r := 0, n-1
	for l < r {
		for l < r && !check(sb[l]) {
			l += 1
		}
		for l < r && !check(sb[r]) {
			r -= 1
		}
		sb[l], sb[r] = sb[r], sb[l]
		l += 1
		r -= 1
	}
	return string(sb)
}

func check(c byte) bool {
	if c == 'a' || c == 'e' || c == 'o' || c == 'i' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'O' || c == 'I' || c == 'U' {
		return true
	}
	return false
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))
}
