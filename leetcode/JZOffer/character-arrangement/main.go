package main

import "fmt"

/**
题目: 输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc，则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
思路: 回溯
*/

func characterArrangement(s []byte, l, r, n int, res *[]string) {

	if l >= n {
		return
	}
	if l == r {
		*res = append(*res, string(s))
		//fmt.Printf("%s\n", s)
		return
	}
	for i := l; i <= r; i++ {
		s[i], s[l] = s[l], s[i]
		//fmt.Printf("before: %s,%d,%d\n", s, i+1, l+1)
		characterArrangement(s, l+1, r, n, res)
		//fmt.Printf("after: %s,%d,%d\n", s, i+1, l+1)
		s[i], s[l] = s[l], s[i]
	}
	return
}

func main() {
	res := make([]string, 0)
	s := "abc"
	characterArrangement([]byte(s), 0, len(s)-1, len(s), &res)
	fmt.Println(res)
	//for _, row := range res {
	//	fmt.Printf("%s\n", row)
	//}
}
