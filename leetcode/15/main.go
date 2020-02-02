package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}
	res := [][]int{}
	tmpMap := map[string]bool{}
	for i, row := range nums {
		l, r := i+1, n-1
		for l < r && l < n && r > i {
			tmpSum := nums[l] + nums[r] + row
			if tmpSum == 0 {
				key := strconv.Itoa(row) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])
				if _, ok := tmpMap[key]; !ok {
					res = append(res, []int{row, nums[l], nums[r]})
					tmpMap[key] = true
				}
				l+=1
				r-=1
			} else if tmpSum > 0 {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{1,-1,-1,0}))
	fmt.Println(threeSum([]int{-2,0,1,1,2}))
}
