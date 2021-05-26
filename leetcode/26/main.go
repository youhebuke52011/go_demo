package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	curValue, curIndex := nums[0], 1
	for i := 1; i < len(nums); i++ {
		for i < len(nums) && nums[i] == curValue {
			i += 1
		}
		if i >= len(nums) {
			break
		}
		curValue = nums[i]
		nums[curIndex], nums[i] = nums[i], nums[curIndex]
		curIndex += 1
	}
	return curIndex
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{}
	//nums := []int{1, 1}
	// nums := []int{1, 1, 1, 1}
	end := removeDuplicates(nums)
	fmt.Println(nums[:end])
}
