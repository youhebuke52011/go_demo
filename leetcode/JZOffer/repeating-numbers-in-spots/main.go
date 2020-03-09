package main

import "fmt"

func RepeatingNumbersInSpots(nums []int) int {
	for i, row := range nums {
		if i != row {
			nums[row], nums[i] = nums[i], nums[row]
		}
	}

	for i, row := range nums {
		if row != i {
			return row
		}
	}
	return -1
}

func main() {
	fmt.Println(RepeatingNumbersInSpots([]int{2, 3, 1, 0, 2, 5, 3}))
}
