package main

import (
	"fmt"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	sb := []byte(s)
	res := 0
	for i := 0; i < len(sb); i++ {
		flag := false
		for i < len(sb) && sb[i] == ' ' {
			i += 1
		}
		for i < len(sb) && check(sb[i]) {
			flag = true
			i += 1
		}
		if flag {
			res +=1
		}
	}
	return res
}

func check(ch byte) bool {
	//if ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '\'' || ch == '-' {
	if ch != ' ' {
		return true
	}
	return false
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments(""))
	fmt.Println(countSegments("       "))
	fmt.Println(countSegments(" a   b !"))
	fmt.Println(countSegments("love live! mu'sic forever"))
}
