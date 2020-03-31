package main

import "fmt"

/**
两数之和
 */

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1
	for l < r {
		tmp := numbers[l] + numbers[r]
		if tmp == target {
			return []int{l+1, r+1}
		} else if tmp > target {
			r -= 1
		} else {
			l += 1
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
