package main

import "fmt"


func TwoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums))
	for i, row := range nums {
		if tmp, ok := tmpMap[target-row]; ok {
			return []int{tmp, i}
		}
		tmpMap[row] = i

	}
	return []int{}
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 17))
}
