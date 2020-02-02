package main

import "fmt"

var tmpMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := []string{}
	n := len(digits)
	if n != 0 {
		solve(0, n, digits, make([]byte, n), &res)
	}
	return res
}

func solve(index, n int, digits string, sb []byte, res *[]string) {
	if index == n {
		*res = append(*res, string(sb))
		return
	}

	for j := 0; j < len(tmpMap[digits[index]]); j++ {
		//tmp := int((digits[index] - '2') * 3)
		//sb[index] = byte('a' + (tmp + j))
		sb[index] = tmpMap[digits[index]][j]
		solve(index+1, n, digits, sb, res)
	}
	//for i := index; i < n; i++ {
	//	tmp := int((digits[i] - '2') * 3)
	//	for j := 0; j < 3; j++ {
	//		sb[index] = byte('a' + (tmp + j))
	//		solve(index+1, n, digits, sb, res)
	//	}
	//}
}

func main() {
	//fmt.Println(letterCombinations(""))
	//fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("23"))
	//fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations("7"))
}
