package main

import "fmt"

/**
出栈序列是否合法
 */

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, len(pushed))
	index := -1
	for in, out := 0, 0; in < len(pushed) && out < len(popped); in++ {
		if pushed[in] != popped[out] {
			index++
			stack[index] = pushed[in]
		} else {
			out++
			for index >= 0 && stack[index] == popped[out] {
				index--
				out++
			}
		}
	}
	return index == -1
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
