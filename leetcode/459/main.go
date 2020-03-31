package main

import "fmt"

/**
判断它是否可以由它的一个子串重复多次构成
 */

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 || n == 1 {
		return false
	}
	for i := 1; i <= n/2; i++ {
		flag := true
		if n%i != 0 {
			continue
		}
		sub := s[:i]
		for j := i; j <= n-i; j += i {
			if s[j:j+i] != sub {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("abcab"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("abac"))
	fmt.Println(repeatedSubstringPattern("ababba"))
	fmt.Println(repeatedSubstringPattern("ababab"))
	fmt.Println(repeatedSubstringPattern("abcd"))
}
