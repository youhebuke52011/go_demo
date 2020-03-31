package main

import "fmt"

/**
判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值至多为 k。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	hsMap := make(map[int]int, length)

	for i, row := range nums {
		if l, ok := hsMap[row]; ok {
			if i-l <= k {
				return true
			}
		}
		hsMap[row] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
