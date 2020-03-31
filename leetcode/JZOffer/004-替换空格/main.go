package main

import "fmt"

// 请实现一个函数，将一个字符串中的空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
// 剑指offer 原题!!! (替换空格)
// 面试只想到 时间复杂度为O(n^2)
// 面试官要O(n) 复杂度
// 思路: 双指针   一个指针(l)放在原byte数组后面, 另外一个(r)放替换后数组后面,如果没遇到空格 l,r一起向左移动,并swap(l,r)
// 如果遇到空格 l+1,r+3,并把替换字符set到r后面三个位置,不满足条件l<=r && l>=0 && r>=0后 return 结果
func bigo(s []byte) {
	n := len(s)
	l, r := n-1, n-1
	for i := n - 1; i >= 0; i-- {
		if s[i] != 0 {
			l = i
			break
		}
	}

	for l < r && l >= 0 && r >= 0 {
		if s[l] == ' ' {
			l-=1
			r-=3
			s[r+1] = '%'
			s[r+2] = '2'
			s[r+3] = '0'
		} else {
			s[l],s[r] = s[r],s[l]
			l-=1
			r-=1
		}
	}
}

func main() {
	str := "We Are Happy"
	s := make([]byte, len(str)+4)
	for i, row := range str {
		s[i] = byte(row)
	}
	bigo(s)
	fmt.Printf("%s",s)
}
