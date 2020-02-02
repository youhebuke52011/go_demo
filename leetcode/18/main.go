package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l, r := j+1, n-1
			for l < r && l < n && r < n {
				tmp := nums[i] + nums[j] + nums[l] + nums[r]
				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					if nums[l+1] == nums[l] {
						l += 1
					} else if nums[r-1] == nums[r] {
						r -= 1
					} else {
						l += 1
						r -= 1
					}
				} else if tmp < target {
					l += 1
				} else {
					r -= 1
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
