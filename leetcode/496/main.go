package main

import "fmt"

/**
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
 */

func nextGreaterElement(findNums []int, nums []int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, row := range nums {
		hashMap[row] = i
	}
	res := make([]int, len(findNums))
	for i, row := range findNums {
		res[i] = -1
		var hm int
		var ok bool
		if hm, ok = hashMap[row]; !ok {
			continue
		}
		for j := hm + 1; j < len(nums); j++ {
			if nums[j] > row {
				res[i] = nums[j]
				break
			}
		}
	}
	return res
}

func main() {
	//fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	//fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{6, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{6, 3, 4, 2}))
}
