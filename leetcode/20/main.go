package main

import (
	"container/list"
	"fmt"
)

// func isValid(s string) bool {
// 	var stack []byte
// 	n := len(s)
// 	if n%2 != 0 {
// 		return false
// 	}
// 	if n == 0 {
// 		return true
// 	} else {
// 		if s[0] == ']' || s[0] == '}' || s[0] == ')' {
// 			return false
// 		} else {
// 			stack = make([]byte, n)
// 			stack[0] = s[0]
// 		}
// 	}
// 	index := 1
// 	for i := 1; len(stack) > 0; i++ {
// 		if i == n && index == 0 {
// 			break
// 		}
// 		if i == n {
// 			return false
// 		}
// 		if s[i] == ']' && index-1 >= 0 && stack[index-1] == '[' {
// 			index -= 1
// 		} else if s[i] == '}' && index-1 >= 0 && stack[index-1] == '{' {
// 			index -= 1
// 		} else if s[i] == ')' && index-1 >= 0 && stack[index-1] == '(' {
// 			index -= 1
// 		} else {
// 			stack[index] = s[i]
// 			index += 1
// 		}
// 	}
// 	return true
// }

func isValid(s string) bool {
	stack := list.New()
	sb := []byte(s)
	for _, row := range sb {
		cur := fmt.Sprintf("%c", row)
		if stack.Len() == 0 && (cur == "]" || cur == ")" || cur == "}") {
			return false
		}
		// if stack.Len() != 0 {
		// 	fmt.Println(cur, stack.Back().Value.(string))
		// }
		if stack.Len() != 0 && (cur == "]" && stack.Back().Value.(string) == "[" || cur == ")" && stack.Back().Value.(string) == "(" || cur == "}" && stack.Back().Value.(string) == "{") {
			stack.Remove(stack.Back())
			continue
		}
		stack.PushBack(cur)
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func main() {
	//res := isValid("()")
	//res := isValid("()[]{}")
	//res := isValid("(]")
	// res := isValid("([)]")
	// res := isValid("(()(")
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("(()("))
}
