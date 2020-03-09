package main

import "fmt"

func removeElement(nums []int, val int) int {
	curIndex := 0
	for i := 0; i < len(nums); i++ {
		for i < len(nums) && val == nums[i] {
			i += 1
		}
		if i >= len(nums) {
			break
		}
		nums[i], nums[curIndex] = nums[curIndex], nums[i]
		curIndex +=1
	}
	return curIndex
}

func main() {
	//nums := []int{3,2,2,3}
	nums := []int{0,1,2,2,3,0,4,2}
	end := removeElement(nums, 2)
	fmt.Println(nums[:end])
}
