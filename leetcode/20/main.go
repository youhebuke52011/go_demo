package main

import "fmt"

func isValid(s string) bool {
	var stack []byte
	n := len(s)
	if n%2 != 0 {
		return false
	}
	if n == 0 {
		return true
	} else {
		if s[0] == ']' || s[0] == '}' || s[0] == ')' {
			return false
		} else {
			stack = make([]byte, n)
			stack[0] = s[0]
		}
	}
	index := 1
	for i := 1; len(stack) > 0; i++ {
		if i == n && index == 0 {
			break
		}
		if i == n {
			return false
		}
		if s[i] == ']' && index-1 >= 0 && stack[index-1] == '[' {
			index -= 1
		} else if s[i] == '}' && index-1 >= 0 && stack[index-1] == '{' {
			index -= 1
		} else if s[i] == ')' && index-1 >= 0 && stack[index-1] == '(' {
			index -= 1
		} else {
			stack[index] = s[i]
			index += 1
		}
	}
	return true
}

func main() {
	//res := isValid("()")
	//res := isValid("()[]{}")
	//res := isValid("(]")
	//res := isValid("([)]")
	//res := isValid("{[]}")
	res := isValid("(()(")
	fmt.Println(res)
}
