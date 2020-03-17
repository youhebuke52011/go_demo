package main

import "fmt"

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
