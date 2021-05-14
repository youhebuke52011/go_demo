package main

import (
	"fmt"
	"sort"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	m := len(nums1)
	if m == 0 {
		return 0
	}
	// size := m * m
	// if m == 1 {
	// 	size = 1
	// }
	hm := make([]int, m*m)
	// fmt.Println(hm)
	for i, row := range nums1 {
		for j, val := range nums2 {
			hm[i*m+j] = row + val
		}
	}
	n := len(hm)

	sort.Ints(hm)
	// fmt.Println(hm)
	var f1, f2 func(int) int
	f1 = func(target int) int {
		// 找最后一个
		res := -1
		l, r := 0, n-1

		for l <= r {
			mid := l + (r-l)/2
			if hm[mid] == target {
				l = mid + 1
				res = mid
			} else if hm[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		return res
	}

	f2 = func(target int) int {
		// 找最开始的那个
		res := -1
		l, r := 0, n-1
		for l <= r {
			mid := l + (r-l)/2
			if hm[mid] == target {
				r = mid - 1
				res = mid
			} else if hm[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return res
	}

	for _, row := range nums3 {
		for _, val := range nums4 {
			target := -(row + val)
			end := f1(target)
			if end != -1 {
				first := f2(target)
				// fmt.Println(end, first, target)
				ans += end - first + 1
			}
		}
	}
	return ans
}

func main() {
	/*
			[0,1,-1]
		[-1,1,0]
		[0,0,1]
		[-1,1,1]
	*/
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	fmt.Println(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	fmt.Println(fourSumCount([]int{0, 1, -1}, []int{-1, 1, 0}, []int{0, 0, 1}, []int{-1, 1, 1}))
}
