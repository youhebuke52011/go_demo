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
	var (
		pre = nums[0]
		cur = max(nums[0], nums[1])
		tmp int
		ret int
	)
	for i := 2; i < n-1; i++ {
		tmp = cur
		cur = max(pre+nums[i], cur)
		pre = tmp
	}
	// fmt.Println(cur)
	ret = cur

	tmp = 0
	pre = nums[n-1]
	cur = max(nums[n-1], nums[n-2])
	for i := n - 3; i > 0; i-- {
		tmp = cur
		cur = max(pre+nums[i], cur)
		pre = tmp
	}
	// fmt.Println(cur)
	ret = max(ret, cur)
	return ret
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}
