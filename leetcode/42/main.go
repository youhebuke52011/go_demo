package main

import "fmt"

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
	left, right := make([]int, length), make([]int, length)

	left[0], right[length-1] = height[0], height[length-1]
	for i := 1; i < length; i++ {
		// height[:i+1] 左边最长的柱子
		left[i] = Max(left[i-1], height[i])
		// height[i:] 右边最长的柱子
		right[length-1-i] = Max(right[length-i], height[length-1-i])
	}

	res := 0
	for i := 0; i < length; i++ {
		res += Min(left[i], right[i]) - height[i]
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
