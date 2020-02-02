package main

import "fmt"

func solve(leftNum, rightNum, index int, s []byte, res *[]string) {
	// 左右括号都用完
	if leftNum == 0 && rightNum == 0 {
		*res = append(*res, string(s))
		return
	}

	// 还有左括号
	if leftNum > 0 {
		//leftNum -= 1
		//s = append(s, '(')
		s[index] = '('
		solve(leftNum-1, rightNum, index+1, s, res)
	}

	if rightNum > 0 && leftNum < rightNum {
		//rightNum -= 1
		//s = append(s, ')')
		s[index] = ')'
		solve(leftNum, rightNum-1, index+1, s, res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	solve(n, n, 0, make([]byte, n*2), &res)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
