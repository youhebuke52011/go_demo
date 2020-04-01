package main

import "fmt"

/**
题目：左旋转字符串，汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！
思路：先翻转a[0:n-1] 再翻转a[n:len(s)-1],最后把整个翻转
*/

func LeftRotationString(s []byte, n int) []byte {
	if n < 0 || n > len(s) {
		return s
	}
	reverseString(s, 0, n-1)
	reverseString(s, n, len(s)-1)
	reverseString(s, 0, len(s)-1)
	return s
}

func reverseString(s []byte, l, r int) {
	if l > r {
		return
	}
	for l < r {
		s[l], s[r] = s[r], s[l]
		l += 1
		r -= 1
	}
}

func main() {
	s := "abcdefg"
	n := 3
	fmt.Println(s)
	res := LeftRotationString([]byte(s), n)
	fmt.Printf("%s", res)
}
