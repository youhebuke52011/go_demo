package main

import (
	"fmt"
	"strings"
)

/**
//给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
//
// 如果不存在最后一个单词，请返回 0 。
//
// 说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
 */

func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")
	for i:=len(split)-1;i>=0;i--  {
		if len(split[i]) != 0 {
			return len(split[i])
		}
	}
	return 0
}

func main() {
	// res := lengthOfLastWord("Hello World")
	res := lengthOfLastWord("a ")
	// res := lengthOfLastWord(" ")
	fmt.Println(res)
}
