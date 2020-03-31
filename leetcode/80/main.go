package main

import "fmt"
/**
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
 */
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
