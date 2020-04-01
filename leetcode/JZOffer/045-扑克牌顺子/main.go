package main

import (
	"fmt"
	"sort"
)

/**
能不能组成顺子，0可以代表所有数
3 5 1 0 4  这个就可以
*/

func poker(nums []int) bool {
	sort.Ints(nums)
	zero := 0
	for zero < len(nums) && nums[zero] == 0 {
		zero++
	}

	for i := zero + 1; i < len(nums); i++ {
		if zero < 0 {
			return false
		}
		if nums[i]-nums[i-1] > 1 {
			zero -= nums[i] - nums[i-1] - 1
		}
	}
	if zero < 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(poker([]int{3, 5, 1, 0, 4}))
	fmt.Println(poker([]int{3, 6, 8, 0, 0}))
}
