package main

import (
	"fmt"
)

//func isHw(s []byte) bool {
//	for l,r := 0,len(s)-1;l<=r;{
//		if s[l]!=s[r] {
//			return false
//		}
//		l+=1
//		r-=1
//	}
//	return true
//}

func lengthOfLongestSubstring(s string) int {
	sByte := []byte(s)
	tmp := map[byte]int{}
	//n := len(s)
	// a[l:r]
	l,r := 0,0
	maxL := 0
	for i,row := range sByte{
		if index,ok := tmp[row];ok {
			// 发现重复
			if  index >= l {
				l = index + 1
			}
		}
		tmp[row] = i
		r+=1
		if maxL < r-l{
			maxL = r-l
		}
		//fmt.Println(l,r,maxL)
	}

	return maxL
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	fmt.Println(lengthOfLongestSubstring("uqinntq"))
}
