package main

import "fmt"

/**
题目：找出有个数字出现次数超过数组一半的，没有则输出0,不能使用额外数组空间
*/

func MoreThanHalfNum(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	tmpValue, tmpCount := nums[0], 1
	for _, row := range nums {
		if tmpCount == 0 {
			tmpValue = row
			tmpCount = 1
		} else if tmpValue == row {
			tmpCount++
		} else if tmpValue != row {
			tmpCount--
		}
	}

	tmpCount = 0
	for _, row := range nums {
		if tmpValue == row {
			tmpCount++
		}
	}
	if tmpCount > length/2 {
		return tmpValue
	}
	return 0

}

func main() {
	fmt.Println(MoreThanHalfNum([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
