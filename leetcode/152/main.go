package main

import (
	"fmt"
)

// func maxProduct(nums []int) int {
// 	res := math.MinInt64
// 	var max func(a, b int) int
// 	max = func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	tmp := 1
// 	for _, row := range nums {
// 		if tmp*row > tmp {
// 			tmp *= row
// 		} else {
// 			tmp = row
// 		}
// 		res = max(res, tmp)
// 	}
// 	return res
// }

// O(n2)
// func maxProduct(nums []int) int {
// 	res := math.MinInt64
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		tmp := 1
// 		for j := i; j < n; j++ {
// 			tmp *= nums[j]
// 			if tmp > res {
// 				res = tmp
// 			}
// 		}
// 	}
// 	return res
// }

func maxProduct(nums []int) int {
	n := len(nums)
	maxdp := make([]int, n)
	mindp := make([]int, n)
	maxdp[0], mindp[0] = nums[0], nums[0]
	var min, max func(int, int) int
	min = func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		}
		return i2
	}
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			maxdp[i] = max(maxdp[i-1]*nums[i], nums[i])
			mindp[i] = min(mindp[i-1]*nums[i], nums[i])
		} else {
			maxdp[i] = max(mindp[i-1]*nums[i], nums[i])
			mindp[i] = min(maxdp[i-1]*nums[i], nums[i])
		}

		if maxdp[i] > res {
			res = maxdp[i]
		}
	}
	return res
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4, 9}))
	fmt.Println(maxProduct([]int{2, 3, -2, 10}))
	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
