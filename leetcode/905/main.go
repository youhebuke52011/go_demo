package main

import "fmt"

/**
偶数放在数组前面
 */

func sortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	for l < r {
		for l < r && A[l]%2 == 0 {
			l += 1
		}
		for l < r && A[r]%2 == 1 {
			r -= 1
		}
		if l < r {
			A[l], A[r] = A[r], A[l]
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParity([]int{1, 2, 3, 4}))
	fmt.Println(sortArrayByParity([]int{6, 2, 1, 4}))
	fmt.Println(sortArrayByParity([]int{3,1,2,4}))
}
