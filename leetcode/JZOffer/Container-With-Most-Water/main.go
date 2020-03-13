package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ContainerWithMostWater(nums []int) int {
	l, r := 0, len(nums)-1
	res := min(nums[l], nums[r]) * (r - l)
	for l < r {
		if nums[l] < nums[r] {
			l++
		} else {
			r--
		}
		res = max(res, min(nums[l], nums[r])*(r-l))
	}
	return res
}

func main() {
	fmt.Println(ContainerWithMostWater([]int{1, 3, 5}))
}
