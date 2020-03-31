package main

import "fmt"

/**
特定规则的括号对数量
 */

func slove(l, r int, sb []byte) int {
	tmp := 0
	res := 0
	for i := l; i <= r; i++ {
		if sb[i] == '(' {
			tmp += 1
		} else if sb[i] == ')' {
			tmp -= 1
		}
		if tmp == 0 {
			if i-l == 1 {
				res += 1
			} else {
				res += 2 * slove(l+1, i, sb)
			}
			l = i + 1
		}

	}
	return res
}

func scoreOfParentheses(S string) int {
	l, r := 0, len(S)-1
	res := slove(l, r, []byte(S))
	return res
}

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("(()(()))"))
}
