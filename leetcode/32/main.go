package main

import (
	"container/list"
	"fmt"
)

func longestValidParentheses(s string) int {
	stack := list.New()
	sb := []byte(s)
	res := 0
	stack.PushBack(-1)
	for i, row := range sb {
		if fmt.Sprintf("%c", row) == ")" {
			stack.Remove(stack.Back())
			if stack.Len() != 0 {
				front := stack.Back().Value.(int)
				if res < i-front {
					res = i - front
					fmt.Println(i, front, res)
				}
				continue
			}
		}
		stack.PushBack(i)
	}
	return res
}

func main() {
	fmt.Println(longestValidParentheses("(())"))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses("()(()"))

}
