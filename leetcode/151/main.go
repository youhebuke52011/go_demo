package main

import (
	"fmt"
	"strings"
)

// 需要借助额外的空间O(n)
//func reverseWords(s string) string {
//	n := len(s)
//	sb := make([]byte, n)
//	//l, r := 0, n-1
//	splits := strings.Split(s, " ")
//	index := 0
//	for j := len(splits) - 1; j >= 0; j-- {
//		row := splits[j]
//		if row == "" {
//			continue
//		}
//		if j != len(splits)-1 && sb[0] != 0 {
//			sb[index] = ' '
//			index++
//		}
//		for i := 0; i < len(row); i++ {
//			//fmt.Println(string(sb))
//			sb[index] = row[i]
//			index++
//		}
//	}
//	test := sb[:index]
//	return string(test)
//}

// 空间复杂度O(1)
func reverseWords(s string) string {

}

func main() {
	//fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  world! hello  "))
}
