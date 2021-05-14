package main

import "fmt"

func validateStackSequences(in []int, out []int) bool {
	stack := make([]int, len(in))
	k := -1
	for i, j := 0, 0; i < len(in); i++ {
		if in[i] != out[j] {
			k++
			stack[k] = in[i]
		} else {
			j++
			for k >= 0 && stack[k] == out[j] {
				k--
				j++
			}
		}
	}
	if k > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	// fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
