package main

import "fmt"

/**
最大子序列和
 */

// dp算法
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := nums[0]
	cur := nums[0]
	for i := 1; i < n; i++ {
		if cur+nums[i] > nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if cur > res {
			res = cur
		}
	}
	return res
}

// 分治算法
//func maxSubArray(nums []int) int {
//
//}


func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
