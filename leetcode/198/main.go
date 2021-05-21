package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// dp := make([]int, n)
	ret := nums[0]
	pre := nums[0]
	cur := max(nums[0], nums[1])
	// dp[0] = nums[0]
	// dp[1] = max(nums[0],nums[1])
	var tmp int
	for i := 2; i < n; i++ {
		tmp = cur
		cur = max(cur, pre+nums[i])
		pre = tmp
		ret = max(ret, cur)
		// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		// ret = max(ret, dp[i])
	}
	// fmt.Println(dp)
	return ret
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 2}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
