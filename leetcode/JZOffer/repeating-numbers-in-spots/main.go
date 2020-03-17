package main

import "fmt"

/**
题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。
也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字2。
（不能借助额外空间）
*/

func RepeatingNumbersInSpots(nums []int) int {
	for _, row := range nums {
		if row < 0 || row > len(nums)-1 {
			return -1
		}
	}

	for i, row := range nums {
		for i != row {
			if row == nums[row] {
				nums[0] = row
				return row
			} else {
				nums[row], nums[i] = nums[i], nums[row]
			}
		}
	}
	return -1
}

func main() {
	//fmt.Println(RepeatingNumbersInSpots([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(RepeatingNumbersInSpots([]int{2, 3, 1, 0, 3, 5, 3}))
}
