package main

import "fmt"

/**
1. 理解 i 点的存水量为 min(max(height[:i+1]...), max(height[i:]...)) - height[i]
1. 高效地找到 max(height[:i+1]...) 和 max(height[i:]...)
 */

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]
	for i := 1; i < length; i++ {
		// height[:i+1] 左边最长的柱子
		left[i] = Max(left[i-1], height[i])
		// height[i:] 右边最长的柱子
		right[length-1-i] = Max(height[length-1-i], right[length-i])
	}

	res := 0
	for i := 0; i < length; i++ {
		res += Min(left[i], right[i]) - height[i]
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	//fmt.Println(trap([]int{}))
}
