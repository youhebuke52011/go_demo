package main

import (
	"fmt"
	"sort"

	//"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n^2)算法
//func lengthOfLIS(nums []int) int {
//	length := len(nums)
//	if length < 2 {
//		return length
//	}
//	// dp[i]标识以nums[i]结尾的最长上升子序列
//	dp := make([]int, length)
//	res := 0
//	for i := 0; i < length; i++ {
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//		res = max(res, dp[i])
//	}
//	return res + 1
//}

// O(nlogn)算法
func lengthOfLIS(nums []int) int {
	length := len(nums)
	res := make([]int, 0, length)
	for _, row := range nums {
		// 搜索 res 中值为 row 的索引，如果找不到，则返回最接近且大于 row 的值的索引，有可能返回len(res)
		index := sort.SearchInts(res, row)
		// 找不到而且没有最接近大于row的
		if index == len(res) {
			res = append(res, row)
		} else if res[index] > row {
			res[index] = row
		}
	}
	return len(res)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
