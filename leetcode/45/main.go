package main

import "fmt"

/**
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
思路: 想清楚，i 定了以后，要保证 next i 为可以 jump 到最远的 i
两个指针,一个记录当前,一个记录之前
*/

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func jump(nums []int) int {
	length := len(nums)
	res := 0
	cur, per, i := 0, 0, 0
	for cur < length-1 {
		per = cur
		for ; i <= per; i++ {
			cur = Max(cur, i+nums[i])
		}
		if per == cur {
			return -1
		}
		res += 1
	}
	return res
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{3,2,1,0,4}))
}
