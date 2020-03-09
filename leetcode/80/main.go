package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	i := 2
	for j := i; j < len(nums); j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	//fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	fmt.Println(removeDuplicates([]int{1,2}))
	//fmt.Println(removeDuplicates([]int{1,2,3}))
}
