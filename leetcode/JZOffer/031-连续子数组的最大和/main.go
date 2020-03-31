package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSubset(nums []int) int {
	res := 0
	tmp := 0
	for _, row := range nums {
		tmp = max(row, tmp+row)
		res = max(res, tmp)
	}
	return res
}

func main() {
	// 8
	fmt.Println(MaxSubset([]int{6, -3, -2, 7, -15, 1, 2, 2}))
}
