package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}
	res, nearest := 0, float64(99999)
	for i, row := range nums {
		l, r := i+1, n-1
		for l < r && r > i && l < n {
			tmpSum := row + nums[l] + nums[r]
			// 有个更靠近的
			if tmpSum > target {
				r -= 1
			} else if tmpSum < target {
				l += 1
			} else {
				return target
			}
			if math.Abs(float64(target-tmpSum)) < nearest {
				nearest = math.Abs(float64(target - tmpSum))
				res = tmpSum
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	// -1 -1 1 1 3    -1
	// -1 3 -1 = 1,  -1 3 1 = 3
	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
}
